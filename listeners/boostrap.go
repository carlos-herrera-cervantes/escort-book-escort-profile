package listeners

import (
	"context"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/singleton"
)

func BootstrapListeners() {
	listener := ProfileStatusListener{
		ProfileStatusRepository: &repositories.ProfileStatusRepository{
			Data: singleton.NewPostgresClient(),
		},
		ProfileStatusCategoryRepository: &repositories.ProfileStatusCategoryRepository{
			Data: singleton.NewPostgresClient(),
		},
	}
	emitterService := services.EmitterService{
		Emitter: singleton.NewEmitter(),
	}

	createProfileStatusListener := make(chan interface{})
	emitterService.AddListener("create.profile.status", createProfileStatusListener)

	listener.HandleCreateProfile(context.Background(), createProfileStatusListener)
}
