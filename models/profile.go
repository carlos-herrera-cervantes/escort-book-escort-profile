package models

import (
	"escort-book-escort-profile/enums"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Profile struct {
	Id            string    `json:"id"`
	EscortId      string    `json:"escortId" validate:"required"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email" validate:"required"`
	PhoneNumber   string    `json:"phoneNumber" validate:"required"`
	Gender        string    `json:"gender" validate:"validateRoles"`
	NationalityId string    `json:"nationalityId"`
	Birthdate     string    `json:"birthdate"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type PartialProfile struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Gender        string `json:"gender"`
	NationalityId string `json:"nationalityId"`
	Birthdate     string `json:"birthdate"`
}

func (p *PartialProfile) MapPartial(profile *Profile) {
	if p.FirstName != "" {
		profile.FirstName = p.FirstName
	}

	if p.LastName != "" {
		profile.LastName = p.LastName
	}

	if p.Gender != "" {
		profile.Gender = p.Gender
	}

	if p.NationalityId != "" {
		profile.NationalityId = p.NationalityId
	}

	if p.Birthdate != "" {
		profile.Birthdate = p.Birthdate
	}
}

func (p *Profile) SetDefaultValues() *Profile {
	p.Id = uuid.NewString()

	if len(p.Gender) == 0 {
		p.Gender = enums.NotSpecified
	}

	return p
}

func validateGenderAttribute(fl validator.FieldLevel) bool {
	gender := fl.Field().String()
	validGenders := map[string]bool{
		enums.Male:         true,
		enums.Female:       true,
		enums.NotSpecified: true,
	}

	return validGenders[gender]
}

func (p *Profile) Validate() error {
	var structValidator = validator.New()

	structValidator.RegisterValidation("validateRoles", validateGenderAttribute)
	structError := structValidator.Struct(p)

	if structError != nil {
		return structError
	}

	return nil
}
