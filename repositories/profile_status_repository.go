package repositories

import (
	"context"
	"time"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type ProfileStatusRepository struct {
	Data *singleton.PostgresClient
}

func (r *ProfileStatusRepository) GetOne(ctx context.Context, id string) (models.ProfileStatus, error) {
	query := `SELECT a.*, b.name
	          FROM profile_status a
			  JOIN profile_status_category b
			  ON a.profile_status_category_id = b.id
			  WHERE escort_id = $1;`
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	profileStatus := models.ProfileStatus{}
	err := row.Scan(
		&profileStatus.Id,
		&profileStatus.ProfileId,
		&profileStatus.ProfileStatusCategoryId,
		&profileStatus.CreatedAt,
		&profileStatus.UpdatedAt,
		&profileStatus.Name,
	)

	if err != nil {
		return profileStatus, err
	}

	return profileStatus, nil
}

func (r *ProfileStatusRepository) Create(ctx context.Context, profileStatus *models.ProfileStatus) error {
	query := "INSERT INTO profile_status VALUES ($1, $2, $3, $4, $5)"
	profileStatus.SetDefaultValues()

	_, err := r.Data.EscortProfileDB.ExecContext(
		ctx,
		query,
		profileStatus.Id,
		profileStatus.ProfileId,
		profileStatus.ProfileStatusCategoryId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *ProfileStatusRepository) UpdateOne(ctx context.Context, id string, profileStatus *models.ProfileStatus) error {
	query := "UPDATE profile_status SET profile_status_category_id = $1, updated_at = $2 WHERE escort_id = $3;"

	_, err := r.Data.EscortProfileDB.ExecContext(
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
