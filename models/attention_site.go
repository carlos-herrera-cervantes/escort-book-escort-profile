package models

import "time"

type AttentionSite struct {
	Id                      string    `json:"id"`
	ProfileId               string    `json:"profileId"`
	AttentionSiteCategoryId string    `json:"attentionSiteCategoryId"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}
