package models

import "time"

type TimeCategory struct {
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	MeasurementUnit string    `json:"measurementUnit"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
