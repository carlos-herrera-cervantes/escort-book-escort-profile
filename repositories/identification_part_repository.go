package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
)

type IdentificationPartRepository struct {
	Data *db.Data
}

func (r *IdentificationPartRepository) GetAll(
	ctx context.Context, offset, limit int,
) ([]models.IdentificationPart, error) {
	query := "SELECT * FROM identification_part OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var parts []models.IdentificationPart

	for rows.Next() {
		var part models.IdentificationPart

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
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var part models.IdentificationPart
	err := row.Scan(&part.Id, &part.Name)

	if err != nil {
		return models.IdentificationPart{}, err
	}

	return part, nil
}

func (r *IdentificationPartRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM identification_part;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
