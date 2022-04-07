package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Service struct {
	Id                string          `json:"id"`
	ProfileId         string          `json:"profileId" validate:"required"`
	ServiceCategoryId string          `json:"serviceCategoryId" validate:"required"`
	Cost              decimal.Decimal `json:"cost" validate:"gte=1"`
	CreatedAt         time.Time       `json:"createdAt"`
	UpdatedAt         time.Time       `json:"updatedAt"`
	Name              string          `json:"name"`
}

func (s *Service) SetDefaultValues() { s.Id = uuid.NewString() }

func (s *Service) Validate() error {
	var structValidator = validator.New()
	structError := structValidator.Struct(s)

	if structError != nil {
		return structError
	}

	return nil
}
