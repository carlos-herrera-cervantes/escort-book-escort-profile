package services

import (
	"mime/multipart"
)

//go:generate mockgen -destination=./mocks/is3_service.go -package=mocks --build_flags=--mod=mod . IS3Service
type IS3Service interface {
	Upload(bucket, filename, profileId string, body multipart.File) (string, error)
}
