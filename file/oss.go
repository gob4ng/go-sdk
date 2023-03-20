package file

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type OssConfig struct {
	HostName        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	Directory       string
	StsData         StsData
}

type StsData struct {
	Region          string
	Scheme          string
	RoleArn         string
	RoleSessionName string
}

type StsCredential struct {
	SecurityToken string
	ExpiredTime   int64
}

type ResponseSts struct {
	RequestID       string `json:"RequestId"`
	AssumedRoleUser struct {
		Arn           string `json:"Arn"`
		AssumedRoleID string `json:"AssumedRoleId"`
	} `json:"AssumedRoleUser"`
	Credentials struct {
		SecurityToken   string    `json:"SecurityToken"`
		AccessKeyID     string    `json:"AccessKeyId"`
		AccessKeySecret string    `json:"AccessKeySecret"`
		Expiration      time.Time `json:"Expiration"`
	} `json:"Credentials"`
}

type OssFunction interface {
	GetSts(data StsData, expiredTime int64) (*StsCredential, *error)
	GetUrlImage(fileName string, credential *StsCredential) (*string, *error)
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

func (o OssConfig) GetSts(data StsData, expiredTime int64) (*StsCredential, *error) {

	stsClient, err := sts.NewClientWithAccessKey(data.Region, o.AccessKeyID, o.AccessKeySecret)
	if err != nil {
		return nil, &err
	}

	request := sts.CreateAssumeRoleRequest()
	request.Scheme = data.Scheme
	request.RoleArn = data.RoleArn
	request.RoleSessionName = data.RoleSessionName

	assumeResponse, err := stsClient.AssumeRole(request)
	if err != nil {
		return nil, &err
	}

	var responseSts ResponseSts
	if err := json.Unmarshal(assumeResponse.GetHttpContentBytes(), &responseSts); err != nil {
		return nil, &err
	}

	stsCredential := StsCredential{
		SecurityToken: responseSts.Credentials.SecurityToken,
		ExpiredTime:   expiredTime,
	}

	return &stsCredential, nil

}

func (o OssConfig) GetUrlImage(fileName string, credential *StsCredential) (*string, *error) {

	securityToken := ""
	var expiredTime int64 = 0
	if credential != nil {
		securityToken = credential.SecurityToken
		expiredTime = credential.ExpiredTime
	}

	client, err := oss.New(o.HostName, o.AccessKeyID, o.AccessKeySecret, oss.SecurityToken(securityToken))
	if err != nil {
		return nil, &err
	}

	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		return nil, &err
	}

	signedURL, err := bucket.SignURL(filepath.Join(o.Directory, fileName), oss.HTTPGet, expiredTime)
	if err != nil {
		return nil, &err
	}

	httpsUrl := strings.Replace(signedURL, "http", "https", 1)

	return &httpsUrl, nil
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

	if err := bucket.GetObjectToFile(filepath.Join(o.Directory, fileName),
		filepath.Join(targetDownload, fileName)); err != nil {
		return nil, &err
	}

	path := filepath.Join(targetDownload, fileName)

	return &path, nil
}
