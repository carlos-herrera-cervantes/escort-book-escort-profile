package singleton

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestNewS3Client(t *testing.T) {
	t.Run("Should return a pointer to S3", func(t *testing.T) {
		s3Client := NewS3Client()
		assert.IsType(t, &s3.S3{}, s3Client)
	})
}
