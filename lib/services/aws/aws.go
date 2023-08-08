package aws

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"mini-project-apotek/constants"
	"mini-project-apotek/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadFileS3(name string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	file.Filename = utils.GenerateRandomString(name)
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return "", err
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(constants.AWS_REGION),
	})
	if err != nil {
		return "", err
	}

	s3Client := s3.New(sess)
	s3Key := fmt.Sprintf("uploads/%s", file.Filename)
	s3Bucket := "apotek"
	objectInput := &s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(s3Key),
		Body:   bytes.NewReader(buf.Bytes()),
	}

	if _, err := s3Client.PutObject(objectInput); err != nil {
		return "", err
	}

	return s3Key, nil
}