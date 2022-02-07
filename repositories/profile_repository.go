package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
)

type ProfileRepository struct {
	Data *db.Data
}

func (r *ProfileRepository) GetOne(ctx context.Context, id string) (models.Profile, error) {
	query := "SELECT * FROM profiles WHERE id = $1"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var profile models.Profile
	err := row.Scan(
		&profile.Id,
		&profile.EscortId,
		&profile.FirstName,
		&profile.LastName,
		&profile.Email,
		&profile.PhoneNumber,
		&profile.Gender,
		&profile.NationalityId,
		&profile.Birthdate,
		&profile.CreatedAt,
		&profile.UpdatedAt)

	if err != nil {
		return models.Profile{}, err
	}

	return profile, nil
}
