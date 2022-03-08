package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
	"time"
)

type AttentionSiteRepository struct {
	Data *db.Data
}

func (r *AttentionSiteRepository) GetAll(
	ctx context.Context,
	profileId string,
	offset,
	limit int,
) ([]models.AttentionSiteDetailed, error) {
	query := `SELECT a.id, a.escort_id, a.attention_site_category_id, a.created_at, a.updated_at, b.name
			  FROM attention_site a
			  join attention_site_category b
			  on a.attention_site_category_id = b.id
			  WHERE escort_id = $3 OFFSET($1) LIMIT($2);`
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit, profileId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sites []models.AttentionSiteDetailed

	for rows.Next() {
		var site models.AttentionSiteDetailed

		rows.Scan(
			&site.Id,
			&site.ProfileId,
			&site.AttentionSiteCategoryId,
			&site.CreatedAt,
			&site.UpdatedAt,
			&site.CategoryName,
		)

		sites = append(sites, site)
	}

	return sites, nil
}

func (r *AttentionSiteRepository) GetOne(ctx context.Context, id string) (models.AttentionSiteDetailed, error) {
	query := `SELECT a.id, a.escort_id, a.attention_site_category_id, a.created_at, a.updated_at, b.name
			  FROM attention_site a
			  join attention_site_category b
			  on a.attention_site_category_id = b.id
			  WHERE escort_id = $1`
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var site models.AttentionSiteDetailed
	err := row.Scan(
		&site.Id,
		&site.ProfileId,
		&site.AttentionSiteCategoryId,
		&site.CreatedAt,
		&site.UpdatedAt)

	if err != nil {
		return models.AttentionSiteDetailed{}, err
	}

	return site, nil
}

func (r *AttentionSiteRepository) Create(ctx context.Context, site *models.AttentionSite) error {
	query := "INSERT INTO attention_site VALUES ($1, $2, $3, $4, $5);"
	site.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		site.Id,
		site.ProfileId,
		site.AttentionSiteCategoryId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *AttentionSiteRepository) DeleteOne(ctx context.Context, id string) error {
	query := "DELETE FROM attention_site WHERE id = $1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *AttentionSiteRepository) Count(ctx context.Context, profileId string) (int, error) {
	query := "SELECT COUNT(id) FROM attention_site WHERE escort_id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, profileId)

	var number int

	row.Scan(&number)

	return number, nil
}
