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

func (r *AttentionSiteRepository) GetAll(ctx context.Context, offset, limit int) ([]models.AttentionSite, error) {
	query := "SELECT * FROM attentionsites OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sites []models.AttentionSite

	for rows.Next() {
		var site models.AttentionSite

		rows.Scan(
			&site.Id,
			&site.ProfileId,
			&site.AttentionSiteCategoryId,
			&site.CreatedAt,
			&site.UpdatedAt)

		sites = append(sites, site)
	}

	return sites, nil
}

func (r *AttentionSiteRepository) GetOne(ctx context.Context, id string) (models.AttentionSite, error) {
	query := "SELECT * FROM attentionsites WHERE profileid = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var site models.AttentionSite
	err := row.Scan(
		&site.Id,
		&site.ProfileId,
		&site.AttentionSiteCategoryId,
		&site.CreatedAt,
		&site.UpdatedAt)

	if err != nil {
		return models.AttentionSite{}, err
	}

	return site, nil
}

func (r *AttentionSiteRepository) Create(ctx context.Context, site *models.AttentionSite) error {
	query := "INSERT INTO attentionsites VALUES ($1, $2, $3, $4, $5);"
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
	query := "DELETE FROM attentionsites WHERE id = $1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *AttentionSiteRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM attentionsites;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
