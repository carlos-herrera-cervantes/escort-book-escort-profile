package repositories

import (
	"context"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type IdentificationPartRepository struct {
	Data *singleton.PostgresClient
}

func (r *IdentificationPartRepository) GetAll(
	ctx context.Context, offset, limit int,
) ([]models.IdentificationPart, error) {
	query := "SELECT * FROM identification_part OFFSET($1) LIMIT($2);"
	rows, err := r.Data.EscortProfileDB.QueryContext(ctx, query, offset, limit)
	parts := []models.IdentificationPart{}

	if err != nil {
		return parts, err
	}

	defer rows.Close()

	for rows.Next() {
		part := models.IdentificationPart{}

		rows.Scan(
			&part.Id,
			&part.Name,
			&part.CreatedAt,
			&part.UpdatedAt)

		parts = append(parts, part)
	}

	return parts, nil
}

func (r *IdentificationPartRepository) GetById(ctx context.Context, id string) (models.IdentificationPart, error) {
	query := "SELECT id, name FROM identification_part WHERE id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	part := models.IdentificationPart{}
	err := row.Scan(&part.Id, &part.Name)

	if err != nil {
		return part, err
	}

	return part, nil
}

func (r *IdentificationPartRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM identification_part;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
