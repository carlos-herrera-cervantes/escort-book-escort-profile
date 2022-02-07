package models

import "time"

type ProfileStatus struct {
	Id                      string    `json:"id"`
	ProfileId               string    `json:"profileId"`
	ProfileStatusCategoryId string    `json:"profileStatusCategoryId"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}
