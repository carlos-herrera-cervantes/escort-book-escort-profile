package services

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"escort-book-escort-profile/config"
	mockSingleton "escort-book-escort-profile/singleton/mocks"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestS3ServiceUpload(t *testing.T) {
	controller := gomock.NewController(t)
	mockS3 := mockSingleton.NewMockIS3(controller)
	s3Service := S3Service{S3Client: mockS3}

	t.Run("Should return error when PutObject fails", func(t *testing.T) {
		mockS3.
			EXPECT().
			PutObject(gomock.Any()).
			Return(&s3.PutObjectOutput{}, errors.New("dummy error")).
			Times(1)

		filename := "profile.png"
		bucket := config.InitS3().Buckets.EscortProfile
		profileId := "6390bef92d11ad34619160b0"
		image, _ := os.Open("../assets/avatar.png")
		defer image.Close()

		path, err := s3Service.Upload(bucket, filename, profileId, image)

		assert.Error(t, err)
		assert.True(t, len(path) == 0)
	})

	t.Run("Should return path when PutObject succeeds", func(t *testing.T) {
		mockS3.
			EXPECT().
			PutObject(gomock.Any()).
			Return(&s3.PutObjectOutput{}, nil).
			Times(1)

		filename := "profile.png"
		bucket := config.InitS3().Buckets.EscortProfile
		profileId := "6390bef92d11ad34619160b0"
		image, _ := os.Open("../assets/avatar.png")
		defer image.Close()

		path, err := s3Service.Upload(bucket, filename, profileId, image)
		expectedPath := fmt.Sprintf("%s/%s/%s/%s", config.InitS3().Endpoint, bucket, profileId, filename)

		assert.NoError(t, err)
		assert.Equal(t, expectedPath, path)
	})
}
