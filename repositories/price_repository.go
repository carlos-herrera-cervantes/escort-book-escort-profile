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

func (r *PriceRepository) GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.Price, error) {
	query := `SELECT a.id, a.cost, a.escort_id, a.time_category_id, a.created_at, a.updated_at, a.quantity, b.name, b.measurement_unit
		      FROM price a
			  INNER JOIN time_category b
			  ON b.id = a.time_category_id
			  WHERE escort_id = $3 OFFSET($1) LIMIT($2);`
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit, profileId)

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
			&price.UpdatedAt,
			&price.Quantity,
			&price.Category,
			&price.MeasurementUnit,
		)

		prices = append(prices, price)
	}

	return prices, nil
}

func (r *PriceRepository) GetOne(ctx context.Context, id string) (models.Price, error) {
	query := "SELECT * FROM price WHERE id = $1;"
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
	query := "INSERT INTO price VALUES ($1, $2, $3, $4, $5, $6, $7);"
	price.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		price.Id,
		price.Cost,
		price.ProfileId,
		price.TimeCategoryId,
		time.Now().UTC(),
		time.Now().UTC(),
		price.Quantity,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PriceRepository) UpdateOne(ctx context.Context, id string, price *models.Price) error {
	query := "UPDATE price set cost = $1, time_category_id = $2 WHERE id = $3;"

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
	query := "DELETE FROM price WHERE id = $1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *PriceRepository) Count(ctx context.Context, profileId string) (int, error) {
	query := "SELECT COUNT(*) FROM price WHERE escort_id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, profileId)

	var number int

	row.Scan(&number)

	return number, nil
}
