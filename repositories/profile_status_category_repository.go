package repositories

import (
	"context"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type ProfileStatusCategoryRepository struct {
	Data *singleton.PostgresClient
}

func (r *ProfileStatusCategoryRepository) GetAll(ctx context.Context, offset, limit int) ([]models.ProfileStatusCategory, error) {
	query := "SELECT * FROM profile_status_category OFFSET($1) LIMIT($2);"
	rows, err := r.Data.EscortProfileDB.QueryContext(ctx, query, offset, limit)
	categories := []models.ProfileStatusCategory{}

	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		category := models.ProfileStatusCategory{}

		rows.Scan(
			&category.Id,
			&category.Name,
			&category.Active,
			&category.CreatedAt,
			&category.UpdatedAt)

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *ProfileStatusCategoryRepository) GetOneByName(ctx context.Context, name string) (models.ProfileStatusCategory, error) {
	query := "SELECT * FROM profile_status_category WHERE name = $1"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, name)

	category := models.ProfileStatusCategory{}
	err := row.Scan(&category.Id, &category.Name, &category.Active, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *ProfileStatusCategoryRepository) GetById(ctx context.Context, id string) (models.ProfileStatusCategory, error) {
	query := "SELECT id, name FROM profile_status_category WHERE id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	var category models.ProfileStatusCategory
	err := row.Scan(&category.Id, &category.Name)

	if err != nil {
		return models.ProfileStatusCategory{}, err
	}

	return category, nil
}

func (r *ProfileStatusCategoryRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM profile_status_category;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
