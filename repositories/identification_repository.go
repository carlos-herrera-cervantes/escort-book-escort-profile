package repositories

import (
	"context"
	"time"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type IdentificationRepository struct {
	Data *singleton.PostgresClient
}

func (r *IdentificationRepository) GetAll(ctx context.Context, profileId string) ([]models.Identification, error) {
	query := "SELECT * FROM identification WHERE escort_id = $1;"
	rows, err := r.Data.EscortProfileDB.QueryContext(ctx, query, profileId)
	identifications := []models.Identification{}

	if err != nil {
		return identifications, err
	}

	defer rows.Close()

	for rows.Next() {
		identification := models.Identification{}

		rows.Scan(
			&identification.Id,
			&identification.Path,
			&identification.ProfileId,
			&identification.IdentificationPartId,
			&identification.CreatedAt,
			&identification.UpdatedAt)

		identifications = append(identifications, identification)
	}

	return identifications, nil
}

func (r *IdentificationRepository) GetOne(ctx context.Context, id string) (models.Identification, error) {
	query := "SELECT * FROM identification WHERE id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	identification := models.Identification{}
	err := row.Scan(
		&identification.Id,
		&identification.Path,
		&identification.ProfileId,
		&identification.IdentificationPartId,
		&identification.CreatedAt,
		&identification.UpdatedAt)

	if err != nil {
		return models.Identification{}, err
	}

	return identification, nil
}

func (r *IdentificationRepository) Create(ctx context.Context, identification *models.Identification) error {
	query := "INSERT INTO identification VALUES ($1, $2, $3, $4, $5, $6);"
	identification.SetDefaultValues()

	_, err := r.Data.EscortProfileDB.ExecContext(
		ctx,
		query,
		identification.Id,
		identification.Path,
		identification.ProfileId,
		identification.IdentificationPartId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *IdentificationRepository) UpdateOne(
	ctx context.Context, id string, identification *models.Identification,
) error {
	query := "UPDATE identification SET path = $1, updated_at = $2 WHERE id = $3;"
	_, err := r.Data.EscortProfileDB.ExecContext(ctx, query, identification.Path, time.Now().UTC(), id)

	if err != nil {
		return err
	}

	return nil
}

func (r *IdentificationRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM identification;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
