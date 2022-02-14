package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
	"time"
)

type PriceRepository struct {
	Data *db.Data
}

func (r *PriceRepository) GetAll(ctx context.Context, offset, limit int) ([]models.Price, error) {
	query := "SELECT * FROM prices OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var prices []models.Price

	for rows.Next() {
		var price models.Price

		rows.Scan(
			&price.Id,
			&price.Cost,
			&price.ProfileId,
			&price.TimeCategoryId,
			&price.CreatedAt,
			&price.UpdatedAt)

		prices = append(prices, price)
	}

	return prices, nil
}

func (r *PriceRepository) GetOne(ctx context.Context, id string) (models.Price, error) {
	query := "SELECT * FROM prices WHERE id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var price models.Price
	err := row.Scan(
		&price.Id,
		&price.Cost,
		&price.ProfileId,
		&price.TimeCategoryId,
		&price.CreatedAt,
		&price.UpdatedAt)

	if err != nil {
		return models.Price{}, err
	}

	return price, nil
}

func (r *PriceRepository) Create(ctx context.Context, price *models.Price) error {
	query := "INSERT INTO prices VALUES ($1, $2, $3, $4, $5, $6);"
	price.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		price.Id,
		price.Cost,
		price.ProfileId,
		price.TimeCategoryId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *PriceRepository) UpdateOne(ctx context.Context, id string, price *models.Price) error {
	query := "UPDATE prices set cost = $1, timecategoryid = $2 WHERE id = $3;"

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		price.Cost,
		price.TimeCategoryId,
		time.Now().UTC(),
		id)

	if err != nil {
		return err
	}

	return nil
}

func (r *PriceRepository) DeleteOne(ctx context.Context, id string) error {
	query := "DELETE FROM prices WHERE id = $1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *PriceRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM prices;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
