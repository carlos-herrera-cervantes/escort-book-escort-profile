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
	query := "SELECT * FROM identificationparts OFFSET($1) LIMIT($2);"
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

func (r *IdentificationPartRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM identificationparts;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
