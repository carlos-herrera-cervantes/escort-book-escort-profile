package services

import (
	"context"
	"mime/multipart"
)

type IS3Service interface {
	Upload(ctx context.Context, bucket, fileName, profileId string, body multipart.File) (string, error)
}
