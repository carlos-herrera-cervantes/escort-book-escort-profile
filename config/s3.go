package config

import "os"

type s3 struct {
	Region      string
	Endpoint    string
	Credentials credentials
	Buckets     buckets
}

type credentials struct {
	Id     string
	Secret string
}

type buckets struct {
	EscortProfile string
}

var singleS3 *s3

func InitS3() *s3 {
	if singleS3 != nil {
		return singleS3
	}

	lock.Lock()
	defer lock.Unlock()

	singleS3 = &s3{
		Region:   os.Getenv("AWS_REGION"),
		Endpoint: os.Getenv("S3_ENPOINT"),
		Credentials: credentials{
			Id:     "na",
			Secret: "na",
		},
		Buckets: buckets{
			EscortProfile: "escort-book-escort-profile",
		},
	}

	return singleS3
}
