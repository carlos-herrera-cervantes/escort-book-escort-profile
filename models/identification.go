package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Identification struct {
	Id                   string    `json:"id"`
	Path                 string    `json:"path"`
	ProfileId            string    `json:"profileId" validate:"required"`
	IdentificationPartId string    `json:"identificationPartId" validate:"required"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

func (i *Identification) SetDefaultValues() *Identification {
	i.Id = uuid.NewString()
	return i
}

func (i *Identification) Validate() error {
	var structValidator = validator.New()
	structError := structValidator.Struct(i)

	if structError != nil {
		return structError
	}

	return nil
}
