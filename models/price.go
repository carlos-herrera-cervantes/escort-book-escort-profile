package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Price struct {
	Id              string          `json:"id"`
	Cost            decimal.Decimal `json:"cost" validate:"required"`
	ProfileId       string          `json:"profileId" validate:"required"`
	TimeCategoryId  string          `json:"timeCategoryId" validate:"required"`
	Category        string          `json:"category"`
	Quantity        int             `json:"quantity" validate:"required"`
	MeasurementUnit string          `json:"measurementUnit"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
}

type PricePartial struct {
	Cost           decimal.Decimal `json:"cost"`
	TimeCategoryId string          `json:"timeCategoryId"`
	Quantity       int             `json:"quantity"`
}

func (p *PricePartial) MapPartial(price *Price) {
	if p.Cost != price.Cost {
		price.Cost = p.Cost
	}

	if p.TimeCategoryId != "" {
		price.TimeCategoryId = p.TimeCategoryId
	}

	if p.Quantity != 0 {
		price.Quantity = p.Quantity
	}
}

func (p *Price) SetDefaultValues() *Price {
	p.Id = uuid.NewString()
	return p
}

func (p *Price) Validate() error {
	var structValidator = validator.New()
	structError := structValidator.Struct(p)

	if structError != nil {
		return structError
	}

	return nil
}
