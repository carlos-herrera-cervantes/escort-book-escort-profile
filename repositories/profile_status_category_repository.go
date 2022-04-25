package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
)

type ProfileStatusCategoryRepository struct {
	Data *db.Data
}

func (r *ProfileStatusCategoryRepository) GetAll(ctx context.Context, offset, limit int) ([]models.ProfileStatusCategory, error) {
	query := "SELECT * FROM profile_status_category OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.ProfileStatusCategory

	for rows.Next() {
		var category models.ProfileStatusCategory

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
	row := r.Data.DB.QueryRowContext(ctx, query, name)

	var category models.ProfileStatusCategory
	err := row.Scan(&category.Id, &category.Name, &category.Active, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return models.ProfileStatusCategory{}, err
	}

	return category, nil
}

func (r *ProfileStatusCategoryRepository) GetById(ctx context.Context, id string) (models.ProfileStatusCategory, error) {
	query := "SELECT id, name FROM profile_status_category WHERE id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var category models.ProfileStatusCategory
	err := row.Scan(&category.Id, &category.Name)

	if err != nil {
		return models.ProfileStatusCategory{}, err
	}

	return category, nil
}

func (r *ProfileStatusCategoryRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM profile_status_category;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
