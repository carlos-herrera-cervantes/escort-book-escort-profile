package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AttentionSite struct {
	Id                      string    `json:"id"`
	ProfileId               string    `json:"profileId" validate:"required"`
	AttentionSiteCategoryId string    `json:"attentionSiteCategoryId" validate:"required"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type AttentionSiteDetailed struct {
	Id                      string    `json:"id"`
	ProfileId               string    `json:"profileId"`
	AttentionSiteCategoryId string    `json:"attentionSiteCategoryId"`
	CategoryName            string    `json:"categoryName"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type AttentionSiteWrapper struct {
	AttentionSiteCategoryId string `json:"attentionSiteCategoryId"`
	User                    struct {
		Email string   `json:"email"`
		Roles []string `json:"roles"`
		Id    string   `json:"id"`
		Iat   int64    `json:"iat"`
		Exp   int64    `json:"exp"`
	}
}

func (s *AttentionSite) SetDefaultValues() *AttentionSite {
	s.Id = uuid.NewString()
	return s
}

func (s *AttentionSite) Validate() error {
	var structValidator = validator.New()
	structError := structValidator.Struct(s)

	if structError != nil {
		return structError
	}

	return nil
}
