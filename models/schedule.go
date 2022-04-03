package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Schedule struct {
	Id        string    `json:"id"`
	From      string    `json:"from" validate:"required"`
	To        string    `json:"to" validate:"required"`
	ProfileId string    `json:"profileId" validate:"required"`
	DayId     string    `json:"dayId" validate:"required"`
	DayName   string    `json:"day"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *Schedule) SetDefaultValues() *Schedule {
	s.Id = uuid.NewString()
	return s
}

func (s *Schedule) Validate() error {
	var structValidator = validator.New()
	structError := structValidator.Struct(s)

	if structError != nil {
		return structError
	}

	return nil
}
