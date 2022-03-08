package models

import (
	"escort-book-escort-profile/types"
	"time"

	"github.com/google/uuid"
)

type ProfileStatus struct {
	Id                      string    `json:"id"`
	ProfileId               string    `json:"profileId" validate:"required"`
	ProfileStatusCategoryId string    `json:"profileStatusCategoryId" validate:"required"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type ProfileStatusWrapper struct {
	ProfileStatusCategoryId string `json:"profileStatusCategoryId"`
	User                    types.DecodedJwt
}

func (p *ProfileStatus) SetDefaultValues() *ProfileStatus {
	p.Id = uuid.NewString()
	return p
}
