package models

import "time"

type Avatar struct {
	Id        string    `json:"id"`
	Path      string    `json:"path"`
	ProfileId string    `json:"profileId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
