package models

import (
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

func (p *ProfileStatus) SetDefaultValues() *ProfileStatus {
	p.Id = uuid.NewString()
	return p
}
