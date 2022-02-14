package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
)

type AttentionSiteCategoryRepository struct {
	Data *db.Data
}

func (r *AttentionSiteCategoryRepository) GetAll(ctx context.Context, offset, limit int) ([]models.AttentionSiteCategory, error) {
	query := "SELECT * FROM attentionsitecategories OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.AttentionSiteCategory

	for rows.Next() {
		var category models.AttentionSiteCategory

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

func (r *AttentionSiteCategoryRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM attentionsitecategories;"
	rows := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	rows.Scan(&number)

	return number, nil
}
