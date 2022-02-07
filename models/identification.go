package models

import "time"

type Identification struct {
	Id                   string    `json:"id"`
	Path                 string    `json:"path"`
	ProfileId            string    `json:"profileId"`
	IdentificationPartId string    `json:"identificationPartId"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
