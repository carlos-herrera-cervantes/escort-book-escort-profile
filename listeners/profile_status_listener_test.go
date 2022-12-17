package listeners

import (
	"context"
	"errors"
	"testing"

	"escort-book-escort-profile/models"
	mockRepositories "escort-book-escort-profile/repositories/mocks"

	"github.com/golang/mock/gomock"
)

func TestProfileStatusListenerHandlerCreateProfile(t *testing.T) {
	controller := gomock.NewController(t)
	mockProfileStatusRepository := mockRepositories.NewMockIProfileStatusRepository(controller)
	mockProfileStatusCategoryRepository := mockRepositories.NewMockIProfileStatusCategoryRepository(controller)
	profileStatusListener := ProfileStatusListener{
		ProfileStatusRepository:         mockProfileStatusRepository,
		ProfileStatusCategoryRepository: mockProfileStatusCategoryRepository,
	}

	t.Run("Should interrupt the process when an error occurs retrieving category", func(t *testing.T) {
		mockProfileStatusCategoryRepository.
			EXPECT().
			GetOneByName(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatusCategory{}, errors.New("dummy error")).
			Times(1)
		mockProfileStatusRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)

		channel := make(chan interface{})
		profileStatusListener.HandleCreateProfile(context.Background(), channel)
		channel <- models.Profile{}
	})

	t.Run("Should log the error when an error occurs creating profile status", func(t *testing.T) {
		mockProfileStatusCategoryRepository.
			EXPECT().
			GetOneByName(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatusCategory{}, nil).
			Times(1)
		mockProfileStatusRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		channel := make(chan interface{})
		profileStatusListener.HandleCreateProfile(context.Background(), channel)
		channel <- models.Profile{}
	})
}
