package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
	"time"
)

type ProfileStatusRepository struct {
	Data *db.Data
}

func (r *ProfileStatusRepository) GetOne(ctx context.Context, id string) (models.ProfileStatus, error) {
	query := "SELECT * FROM profile_status WHERE escort_id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var profileStatus models.ProfileStatus
	err := row.Scan(
		&profileStatus.Id,
		&profileStatus.ProfileId,
		&profileStatus.ProfileStatusCategoryId,
		&profileStatus.CreatedAt,
		&profileStatus.UpdatedAt)

	if err != nil {
		return models.ProfileStatus{}, err
	}

	return profileStatus, nil
}

func (r *ProfileStatusRepository) Create(ctx context.Context, profileStatus *models.ProfileStatus) error {
	query := "INSERT INTO profile_status VALUES ($1, $2, $3, $4, $5)"
	profileStatus.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		profileStatus.Id,
		profileStatus.ProfileId,
		profileStatus.ProfileStatusCategoryId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return nil
	}

	return nil
}

func (r *ProfileStatusRepository) UpdateOne(ctx context.Context, id string, profileStatus *models.ProfileStatus) error {
	query := "UPDATE profile_status SET profile_status_category_id = $1, updated_at = $2 WHERE escort_id = $3;"

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		profileStatus.ProfileStatusCategoryId,
		time.Now().UTC(),
		id)

	if err != nil {
		return err
	}

	return nil
}
