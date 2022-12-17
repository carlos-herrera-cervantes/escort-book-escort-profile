package repositories

import (
	"context"
	"time"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type BiographyRepository struct {
	Data *singleton.PostgresClient
}

func (r *BiographyRepository) GetOne(ctx context.Context, id string) (models.Biography, error) {
	query := "SELECT * FROM biography WHERE escort_id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	biography := models.Biography{}
	err := row.Scan(
		&biography.Id,
		&biography.Description,
		&biography.ProfileId,
		&biography.CreatedAt,
		&biography.UpdatedAt)

	if err != nil {
		return models.Biography{}, err
	}

	return biography, nil
}

func (r *BiographyRepository) Create(ctx context.Context, biography *models.Biography) error {
	query := "INSERT INTO biography VALUES ($1, $2, $3, $4, $5);"
	biography.SetDefaultValues()

	_, err := r.Data.EscortProfileDB.ExecContext(
		ctx,
		query,
		biography.Id,
		biography.Description,
		biography.ProfileId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *BiographyRepository) UpdateOne(ctx context.Context, id string, biography *models.Biography) error {
	query := "UPDATE biography set description = $1, updated_at = $2 WHERE escort_id = $3;"

	_, err := r.Data.EscortProfileDB.ExecContext(
		ctx,
		query,
		biography.Description,
		time.Now().UTC(),
		id)

	if err != nil {
		return err
	}

	return nil
}

func (r *BiographyRepository) DeleteOne(ctx context.Context, id string) error {
	query := "DELETE FROM biography WHERE escort_id = $1;"
	_, err := r.Data.EscortProfileDB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
