package services

import (
	"fmt"
	"mime/multipart"

	"escort-book-escort-profile/config"
	"escort-book-escort-profile/singleton"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	S3Client singleton.IS3
}

func (s *S3Service) Upload(
	bucket, filename, profileId string, body multipart.File,
) (string, error) {
	_, err := s.S3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", profileId, filename)),
		Body:   aws.ReadSeekCloser(body),
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s/%s", config.InitS3().Endpoint, bucket, profileId, filename), nil
}
