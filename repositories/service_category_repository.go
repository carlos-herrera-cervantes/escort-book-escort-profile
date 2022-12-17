package repositories

import (
	"context"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type ServiceCategoryRepository struct {
	Data *singleton.PostgresClient
}

func (r *ServiceCategoryRepository) GetAll(ctx context.Context, offset, limit int) ([]models.ServiceCategory, error) {
	query := "SELECT * FROM service_category OFFSET($1) LIMIT($2);"
	rows, err := r.Data.EscortProfileDB.QueryContext(ctx, query, offset, limit)
	categories := []models.ServiceCategory{}

	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		category := models.ServiceCategory{}
		rows.Scan(&category.Id, &category.Name, &category.Active, &category.CreatedAt, &category.UpdatedAt)
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *ServiceCategoryRepository) GetById(ctx context.Context, id string) (models.ServiceCategory, error) {
	query := "SELECT id, name FROM service_category WHERE id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	category := models.ServiceCategory{}
	err := row.Scan(&category.Id, &category.Name)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *ServiceCategoryRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM service_category;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
