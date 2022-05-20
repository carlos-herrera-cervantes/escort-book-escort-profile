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
	Name                    string    `json:"name"`
}

type PartialProfileStatus struct {
	ProfileStatusCategoryId string `json:"profileStatusCategoryId"`
}

func (p *ProfileStatus) SetDefaultValues() { p.Id = uuid.NewString() }
