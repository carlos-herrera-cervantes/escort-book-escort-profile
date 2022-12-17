package repositories

import (
	"context"
	"time"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type PhotoRepository struct {
	Data *singleton.PostgresClient
}

func (r *PhotoRepository) GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.Photo, error) {
	query := "SELECT * FROM photo WHERE escort_id = $3 OFFSET($1) LIMIT($2);"
	rows, err := r.Data.EscortProfileDB.QueryContext(ctx, query, offset, limit, profileId)
	photos := []models.Photo{}

	if err != nil {
		return photos, err
	}

	defer rows.Close()

	for rows.Next() {
		photo := models.Photo{}

		rows.Scan(
			&photo.Id,
			&photo.Path,
			&photo.ProfileId,
			&photo.CreatedAt,
			&photo.UpdatedAt)

		photos = append(photos, photo)
	}

	return photos, nil
}

func (r *PhotoRepository) GetOne(ctx context.Context, id string) (models.Photo, error) {
	query := "SELECT * FROM photo WHERE id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	photo := models.Photo{}
	err := row.Scan(
		&photo.Id,
		&photo.Path,
		&photo.ProfileId,
		&photo.CreatedAt,
		&photo.UpdatedAt)

	if err != nil {
		return models.Photo{}, err
	}

	return photo, nil
}

func (r *PhotoRepository) Create(ctx context.Context, photo *models.Photo) error {
	query := "INSERT INTO photo VALUES ($1, $2, $3, $4, $5);"
	photo.SetDefaultValues()

	_, err := r.Data.EscortProfileDB.ExecContext(
		ctx,
		query,
		photo.Id,
		photo.Path,
		photo.ProfileId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *PhotoRepository) DeleteOne(ctx context.Context, id string) error {
	query := "DELETE FROM photo WHERE id = $1;"
	_, err := r.Data.EscortProfileDB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *PhotoRepository) Count(ctx context.Context, profileId string) (int, error) {
	query := "SELECT COUNT(*) FROM photo WHERE escort_id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, profileId)

	var number int

	row.Scan(&number)

	return number, nil
}
