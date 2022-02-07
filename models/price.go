package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Price struct {
	Id        string          `json:"id"`
	Cost      decimal.Decimal `json:"cost"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
