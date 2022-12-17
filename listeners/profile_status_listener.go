package listeners

import (
	"context"
	"fmt"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"

	log "github.com/inconshreveable/log15"
)

var logger = log.New("listeners")

type ProfileStatusListener struct {
	ProfileStatusRepository         repositories.IProfileStatusRepository
	ProfileStatusCategoryRepository repositories.IProfileStatusCategoryRepository
}

func (l *ProfileStatusListener) HandleCreateProfile(ctx context.Context, listener chan interface{}) {
	go func() {
		for {
			event := <-listener
			profile := event.(models.Profile)
			category, err := l.ProfileStatusCategoryRepository.GetOneByName(ctx, "In Review")

			if err != nil {
				log.Error(fmt.Sprintf("ERROR RETRIEVING CATEGORY: %s", err.Error()))
				continue
			}

			profileStatus := models.ProfileStatus{
				ProfileId:               profile.Id,
				ProfileStatusCategoryId: category.Id,
			}

			if err := l.ProfileStatusRepository.Create(ctx, &profileStatus); err != nil {
				log.Error(fmt.Sprintf("LISTENER ERROR: %s", err.Error()))
			}
		}
	}()
}
