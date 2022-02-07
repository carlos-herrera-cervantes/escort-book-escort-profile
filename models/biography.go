package models

import "time"

type Biography struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	ProfileId   string    `json:"profileId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
