package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
)

type TimeCategoryRepository struct {
	Data *db.Data
}

func (r *TimeCategoryRepository) GetAll(ctx context.Context, offset, limit int) ([]models.TimeCategory, error) {
	query := "SELECT * FROM time_category OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.TimeCategory

	for rows.Next() {
		var category models.TimeCategory

		rows.Scan(
			&category.Id,
			&category.Name,
			&category.MeasurementUnit,
			&category.CreatedAt,
			&category.UpdatedAt)

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *TimeCategoryRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM time_category;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
