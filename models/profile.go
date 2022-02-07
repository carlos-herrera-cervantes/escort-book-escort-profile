package models

import "time"

type Profile struct {
	Id            string    `json:"id"`
	EscortId      string    `json:"escortId"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	Gender        string    `json:"gender"`
	NationalityId string    `json:"nationalityId"`
	Birthdate     string    `json:"birthdate"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
