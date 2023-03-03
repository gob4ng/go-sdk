package ftp

import (
	"bufio"
	"errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

type SftpData struct {
	Hostname                    string
	Protocol                    string
	Port                        string
	Username                    string
	Password                    string
	LocalDirectory              string
	RemoteDirectory             string
	Filename                    string
	CustomPublicKeyFileLocation *string
}

type sshClientConfig struct {
	config   ssh.ClientConfig
	sftpData SftpData
}

func NewSftpConfig(data SftpData) (*sshClientConfig, *error) {

	if err := checkPort(data.Port); err != nil {
		return nil, err
	}

	sshPublicKey, err := checkPublicKey(data.Hostname, data.CustomPublicKeyFileLocation)
	if err != nil {
		return nil, err
	}

	authMethod, err := getAuthMethod(data.Password)
	if err != nil {
		return nil, err
	}

	config := ssh.ClientConfig{
		User:            data.Username,
		Auth:            *authMethod,
		HostKeyCallback: ssh.FixedHostKey(*sshPublicKey),
	}

	sshClientConfig := sshClientConfig{
		config:   config,
		sftpData: data,
	}

	return &sshClientConfig, nil

}

func (s *sshClientConfig) Connect() *error {

	_, err := checkConnection(s)
	if err != nil {
		return err
	}

	return nil
}

func (s *sshClientConfig) Upload() (*int64, *error) {

	client, err := checkConnection(s)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	destinationFile, errDestinationFile := client.Create("." +
		filepath.Join(s.sftpData.RemoteDirectory, s.sftpData.Filename))
	if errDestinationFile != nil {
		return nil, &errDestinationFile
	}
	defer destinationFile.Close()

	workingDirectory, errGetWd := os.Getwd()
	if errGetWd != nil {
		return nil, &errGetWd
	}

	sourceFile, errSourceFile := os.Open(filepath.Join(workingDirectory,
		s.sftpData.LocalDirectory, s.sftpData.Filename))
	if errSourceFile != nil {
		return nil, &errSourceFile
	}
	defer sourceFile.Close()

	byteFile, errByteFile := io.Copy(destinationFile, sourceFile)
	if errByteFile != nil {
		return nil, &errByteFile
	}

	return &byteFile, nil
}

func (s *sshClientConfig) Download() (*int64, *error) {

	client, err := checkConnection(s)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	destinationFile, errDestinationFile := os.Create("." +
		filepath.Join(s.sftpData.LocalDirectory, s.sftpData.Filename))
	if errDestinationFile != nil {
		return nil, &errDestinationFile
	}
	defer destinationFile.Close()

	workingDirectory, errGetWd := os.Getwd()
	if errGetWd != nil {
		return nil, &errGetWd
	}

	sourceFile, errSourceFile := client.Open(filepath.Join(workingDirectory,
		s.sftpData.RemoteDirectory, s.sftpData.Filename))
	if errSourceFile != nil {
		return nil, &errSourceFile
	}
	defer sourceFile.Close()

	byteFile, errByteFile := io.Copy(destinationFile, sourceFile)
	if errByteFile != nil {
		return nil, &errByteFile
	}

	if err := destinationFile.Sync(); err != nil {
		return nil, &err
	}

	return &byteFile, nil
}

func checkPort(port string) *error {

	if port == "" {
		newError := errors.New("port is required")
		return &newError
	}

	portNumber, err := strconv.Atoi(port)
	if err != nil {
		return &err
	}

	if portNumber < 1 || portNumber > 65535 {
		strError := "invalid port: %v" + port
		newError := errors.New(strError)
		return &newError
	}

	return nil
}

func checkPublicKey(hostname string, customPublicKeyFileLocation *string) (*ssh.PublicKey, *error) {

	var file fs.File
	if customPublicKeyFileLocation != nil {

		workingDirectory, err := os.Getwd()
		if err != nil {
			return nil, &err
		}

		file, err = os.Open(filepath.Join(workingDirectory, *customPublicKeyFileLocation))
		if err != nil {
			return nil, &err
		}

	} else {

		userDirectory, err := user.Current()
		if err != nil {
			return nil, &err
		}

		file, err = os.Open(filepath.Join(userDirectory.HomeDir, ".ssh", "id_rsa"))
		if err != nil {
			return nil, &err
		}

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], hostname) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, &err
			}
		}
	}

	if hostKey == nil {
		newError := errors.New("pub key not found")
		return nil, &newError
	}

	return &hostKey, nil
}

func getAuthMethod(password string) (*[]ssh.AuthMethod, *error) {

	if password != "" {
		authMethods := []ssh.AuthMethod{ssh.Password(password)}
		return &authMethods, nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return nil, &err
	}

	sshDirectory := filepath.Join(currentUser.HomeDir, ".ssh", "id_rsa")
	key, err := ioutil.ReadFile(sshDirectory)
	if err != nil {
		return nil, &err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, &err
	}

	authMethods := []ssh.AuthMethod{ssh.PublicKeys(signer)}

	return &authMethods, nil
}

func checkConnection(s *sshClientConfig) (*sftp.Client, *error) {

	connection, err := ssh.Dial(s.sftpData.Protocol,
		s.sftpData.Hostname+":"+s.sftpData.Port, &s.config)
	if err != nil {
		return nil, &err
	}
	defer connection.Close()

	client, err := sftp.NewClient(connection)
	if err != nil {
		return nil, &err
	}

	return client, nil
}
