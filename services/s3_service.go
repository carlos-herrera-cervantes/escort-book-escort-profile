package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct{}

var s3Client *s3.S3

func getS3Client() *s3.S3 {
	if s3Client == nil {
		lock.Lock()
		defer lock.Unlock()

		if s3Client == nil {
			sess, _ := session.NewSession(&aws.Config{
				Region:           aws.String(os.Getenv("AWS_REGION")),
				Credentials:      credentials.NewStaticCredentials("na", "na", ""),
				S3ForcePathStyle: aws.Bool(true),
				Endpoint:         aws.String(os.Getenv("S3_ENPOINT")),
			})
			s3Client = s3.New(sess)
		}
	}

	return s3Client
}

func (s *S3Service) Upload(
	ctx context.Context, bucket, fileName, profileId string, body multipart.File,
) (string, error) {
	client := getS3Client()
	_, err := client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", profileId, fileName)),
		Body:   aws.ReadSeekCloser(body),
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s/%s", os.Getenv("S3_ENPOINT"), bucket, profileId, fileName), nil
}
