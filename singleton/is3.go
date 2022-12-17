package singleton

import "github.com/aws/aws-sdk-go/service/s3"

//go:generate mockgen -destination=./mocks/is3.go -package=mocks --build_flags=--mod=mod . IS3
type IS3 interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}
