package listeners

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
)

func BoostrapListeners() {
	listener := ProfileStatusListener{
		ProfileStatusRepository: &repositories.ProfileStatusRepository{
			Data: db.New(),
		},
		ProfileStatusCategoryRepository: &repositories.ProfileStatusCategoryRepository{
			Data: db.New(),
		},
	}
	emitterService := services.EmitterService{}

	createProfileStatusListener := make(chan interface{})
	emitterService.AddListener("create.profile.status", createProfileStatusListener)

	listener.HandleCreateProfile(context.Background(), createProfileStatusListener)
}
