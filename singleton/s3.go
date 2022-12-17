package singleton

import (
	"log"
	"sync"

	"escort-book-escort-profile/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3
var singleS3Client sync.Once

func initS3() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.InitS3().Region),
		Credentials: credentials.NewStaticCredentials(
			config.InitS3().Credentials.Id,
			config.InitS3().Credentials.Secret,
			"",
		),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(config.InitS3().Endpoint),
	})

	if err != nil {
		log.Panic("Error when creating S3 client: ", err.Error())
	}

	s3Client = s3.New(sess)
}

func NewS3Client() *s3.S3 {
	singleS3Client.Do(initS3)
	return s3Client
}
