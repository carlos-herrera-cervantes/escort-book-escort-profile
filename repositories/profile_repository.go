package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
	"time"
)

type ProfileRepository struct {
	Data *db.Data
}

func (r *ProfileRepository) GetOne(ctx context.Context, id string) (models.Profile, error) {
	query := "SELECT * FROM profiles WHERE id = $1;"
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

func (r *ProfileRepository) Create(ctx context.Context, profile *models.Profile) error {
	query := "INSERT INTO profiles VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
	profile.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		profile.Id,
		profile.EscortId,
		profile.FirstName,
		profile.LastName,
		profile.Email,
		profile.PhoneNumber,
		profile.Gender,
		profile.NationalityId,
		profile.Birthdate,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *ProfileRepository) UpdateOne(ctx context.Context, id string, profile *models.Profile) error {
	query := "UPDATE profiles set firstname=$1, lastname=$2, gender=$3, birthdate=$4, updatedat=$5 WHERE id=$6;"

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		profile.FirstName,
		profile.LastName,
		profile.Gender,
		profile.Birthdate,
		time.Now().UTC(),
		id)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProfileRepository) DeleteOne(ctx context.Context, id string) error {
	query := "DELETE FROM profiles WHERE id=$1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
