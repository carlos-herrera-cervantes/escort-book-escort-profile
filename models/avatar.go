package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Avatar struct {
	Id        string    `json:"id"`
	Path      string    `json:"path" validate:"required"`
	ProfileId string    `json:"profileId" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (a *Avatar) SetDefaultValues() *Avatar {
	a.Id = uuid.NewString()
	return a
}

func (a *Avatar) Validate() error {
	var structValidator = validator.New()
	structError := structValidator.Struct(a)

	if structError != nil {
		return structError
	}

	return nil
}
