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
	query := "SELECT * FROM profile WHERE escort_id = $1;"
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
	query := "INSERT INTO profile VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
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
	query := `UPDATE profile set
			  first_name=$1, last_name=$2, gender=$3, birthdate=$4, updated_at=$5, nationality_id=$6
			  WHERE escort_id=$7;`

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		profile.FirstName,
		profile.LastName,
		profile.Gender,
		profile.Birthdate,
		time.Now().UTC(),
		profile.NationalityId,
		id)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProfileRepository) DeleteOne(ctx context.Context, id string) error {
	query := "DELETE FROM profile WHERE escort_id=$1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
