package models

import (
	"escort-book-escort-profile/types"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Biography struct {
	Id          string    `json:"id"`
	Description string    `json:"description" validate:"required"`
	ProfileId   string    `json:"profileId" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type BiographyWrapper struct {
	Description string `json:"description"`
	User        types.DecodedJwt
}

func (b *Biography) SetDefaultValues() *Biography {
	b.Id = uuid.NewString()
	return b
}

func (b *Biography) Validate() error {
	var structValidator = validator.New()
	structError := structValidator.Struct(b)

	if structError != nil {
		return structError
	}

	return nil
}
