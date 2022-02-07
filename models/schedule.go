package models

import "time"

type Schedule struct {
	Id        string    `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	ProfileId string    `json:"profileId"`
	DayId     string    `json:"dayId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
