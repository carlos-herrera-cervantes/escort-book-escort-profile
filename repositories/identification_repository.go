package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
	"time"
)

type IdentificationRepository struct {
	Data *db.Data
}

func (r *IdentificationRepository) GetAll(
	ctx context.Context, profileId string, offset, limit int,
) ([]models.Identification, error) {
	query := "SELECT * FROM identification WHERE escort_id = $3 OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit, profileId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var identifications []models.Identification

	for rows.Next() {
		var identification models.Identification

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
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var identification models.Identification
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

	_, err := r.Data.DB.ExecContext(
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
	_, err := r.Data.DB.ExecContext(ctx, query, identification.Path, time.Now().UTC(), id)

	if err != nil {
		return err
	}

	return nil
}

func (r *IdentificationRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM identification;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
