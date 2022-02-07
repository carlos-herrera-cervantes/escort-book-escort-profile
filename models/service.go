package models

import "time"

type Service struct {
	Id                string    `json:"id"`
	ProfileId         string    `json:"profileId"`
	ServiceCategoryId string    `json:"serviceCategoryId"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
