package file

import (
	"bitbucket.org/dab-ekyc/app/domains/types"
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path/filepath"
)

type OssConfig struct {
	HostName        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	Directory       string
	SecurityToken   string
}

type StsData struct {
	Region          string
	Scheme          string
	RoleArn         string
	RoleSessionName string
	ExpiredTime     int
}

type OssFunction interface {
	GetSts(data StsData) *error
	Upload(directory string, fileName string, destinationFileName string) (*string, *error)
	Download(fileName string, targetDownload string) (*string, *error)
}

func NewOssConfig(config OssConfig) OssFunction {
	return OssConfig{
		HostName:        config.HostName,
		AccessKeyID:     config.AccessKeyID,
		AccessKeySecret: config.AccessKeySecret,
		BucketName:      config.BucketName,
		Directory:       config.Directory,
	}
}

func (o OssConfig) GetSts(data StsData) *error {

	stsClient, err := sts.NewClientWithAccessKey(data.Region, o.AccessKeyID, o.AccessKeySecret)
	if err != nil {
		return &err
	}

	request := sts.CreateAssumeRoleRequest()
	request.Scheme = data.Scheme
	request.RoleArn = data.RoleArn
	request.RoleSessionName = data.RoleSessionName

	assumeResponse, err := stsClient.AssumeRole(request)
	if err != nil {
		return &err
	}

	var responseSts types.ResponseSts
	if err := json.Unmarshal(assumeResponse.GetHttpContentBytes(), &responseSts); err != nil {
		return &err
	}

	o.SecurityToken = responseSts.Credentials.SecurityToken

	return nil

}

func (o OssConfig) Upload(directory string, fileName string, destinationFileName string) (*string, *error) {

	client, err := oss.New(o.HostName, o.AccessKeyID, o.AccessKeySecret)
	if err != nil {
		return nil, &err
	}

	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		return nil, &err
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, &err
	}

	file, err := os.Open(filepath.Join(wd, directory, fileName))
	if err != nil {
		return nil, &err
	}

	if err := bucket.PutObject(filepath.Join(o.Directory, o.Directory), file); err != nil {
		return nil, &err
	}

	fullPath := "https://" + o.BucketName + "." + filepath.Join(o.HostName, directory, destinationFileName)

	return &fullPath, nil
}

func (o OssConfig) Download(fileName string, targetDownload string) (*string, *error) {

	client, err := oss.New(o.HostName, o.AccessKeyID, o.AccessKeySecret)
	if err != nil {
		return nil, &err
	}

	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		return nil, &err
	}

	if err := bucket.GetObjectToFile(filepath.Join(o.Directory, fileName), filepath.Join(targetDownload, fileName)); err != nil {
		return nil, &err
	}

	path := filepath.Join(targetDownload, fileName)

	return &path, nil
}
