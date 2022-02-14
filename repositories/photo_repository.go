package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
	"time"
)

type PhotoRepository struct {
	Data *db.Data
}

func (r *PhotoRepository) GetAll(ctx context.Context, offset, limit int) ([]models.Photo, error) {
	query := "SELECT * FROM photos OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var photos []models.Photo

	for rows.Next() {
		var photo models.Photo

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
	query := "SELECT * FROM photos WHERE id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var photo models.Photo
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
	query := "INSERT INTO photos VALUES ($1, $2, $3, $4, $5);"
	photo.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
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
	query := "DELETE FROM photos WHERE id = $1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *PhotoRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM photos;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
