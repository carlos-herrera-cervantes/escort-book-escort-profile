package listeners

import (
	"context"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"log"
)

type ProfileStatusListener struct {
	ProfileStatusRepository         *repositories.ProfileStatusRepository
	ProfileStatusCategoryRepository *repositories.ProfileStatusCategoryRepository
}

func (l *ProfileStatusListener) HandleCreateProfile(ctx context.Context, listener chan interface{}) {
	go func() {
		for {
			event := <-listener
			profile := event.(models.Profile)
			category, err := l.ProfileStatusCategoryRepository.GetOneByName(ctx, "In Review")

			log.Println("ERROR RETRIEVING CATEGORY: ", err)

			profileStatus := models.ProfileStatus{
				ProfileId:               profile.Id,
				ProfileStatusCategoryId: category.Id,
			}

			if err := l.ProfileStatusRepository.Create(ctx, &profileStatus); err != nil {
				log.Println("LISTENER ERROR: ", err)
			}
		}
	}()
}
