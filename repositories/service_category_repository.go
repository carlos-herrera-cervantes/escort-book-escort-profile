package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
)

type ServiceCategoryRepository struct {
	Data *db.Data
}

func (r *ServiceCategoryRepository) GetAll(ctx context.Context, offset, limit int) ([]models.ServiceCategory, error) {
	query := "SELECT * FROM service_category OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.ServiceCategory

	for rows.Next() {
		var category models.ServiceCategory
		rows.Scan(&category.Id, &category.Name, &category.Active, &category.CreatedAt, &category.UpdatedAt)
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *ServiceCategoryRepository) GetById(ctx context.Context, id string) (models.ServiceCategory, error) {
	query := "SELECT id, name FROM service_category WHERE id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var category models.ServiceCategory
	err := row.Scan(&category.Id, &category.Name)

	if err != nil {
		return models.ServiceCategory{}, err
	}

	return category, nil
}

func (r *ServiceCategoryRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM service_category;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
