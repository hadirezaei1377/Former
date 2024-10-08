package services

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage struct {
	Bucket string
	Client *s3.S3
}

func NewS3Storage(bucket string, region string) *S3Storage {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return &S3Storage{
		Bucket: bucket,
		Client: s3.New(sess),
	}
}

func (s *S3Storage) Upload(file multipart.File, fileName string) error {
	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, file)
	if err != nil {
		return err
	}

	_, err = s.Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buffer.Bytes()),
	})
	return err
}

func (s *S3Storage) Download(fileName string) ([]byte, error) {
	result, err := s.Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return nil, err
	}

	defer result.Body.Close()
	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, result.Body)
	return buffer.Bytes(), err
}
