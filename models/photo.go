package models

import (
	"escort-book-escort-profile/types"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Photo struct {
	Id        string    `json:"id"`
	Path      string    `json:"path" validate:"required"`
	ProfileId string    `json:"profileId" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PhotoWrapper struct {
	Path string `json:"path"`
	User types.DecodedJwt
}

func (p *Photo) SetDefaultValues() *Photo {
	p.Id = uuid.NewString()
	return p
}

func (p *Photo) Validate() error {
	var structValidator = validator.New()
	structError := structValidator.Struct(p)

	if structError != nil {
		return structError
	}

	return nil
}
