package repositories

import (
	"context"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type TimeCategoryRepository struct {
	Data *singleton.PostgresClient
}

func (r *TimeCategoryRepository) GetAll(ctx context.Context, offset, limit int) ([]models.TimeCategory, error) {
	query := "SELECT * FROM time_category OFFSET($1) LIMIT($2);"
	rows, err := r.Data.EscortProfileDB.QueryContext(ctx, query, offset, limit)
	categories := []models.TimeCategory{}

	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		category := models.TimeCategory{}

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

func (r *TimeCategoryRepository) GetById(ctx context.Context, id string) (models.TimeCategory, error) {
	query := "SELECT id, name, measurement_unit from time_category WHERE id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	category := models.TimeCategory{}
	err := row.Scan(&category.Id, &category.Name, &category.MeasurementUnit)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *TimeCategoryRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM time_category;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
