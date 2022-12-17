package repositories

import (
	"context"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type DayRepository struct {
	Data *singleton.PostgresClient
}

func (r *DayRepository) GetAll(ctx context.Context, offset, limit int) ([]models.Day, error) {
	query := "SELECT * FROM day OFFSET($1) LIMIT($2);"
	rows, err := r.Data.EscortProfileDB.QueryContext(ctx, query, offset, limit)
	days := []models.Day{}

	if err != nil {
		return days, err
	}

	defer rows.Close()

	for rows.Next() {
		day := models.Day{}

		rows.Scan(
			&day.Id,
			&day.Name,
			&day.Active,
			&day.CreatedAt,
			&day.UpdatedAt)

		days = append(days, day)
	}

	return days, nil
}

func (r *DayRepository) GetOneByName(ctx context.Context, name string) (models.Day, error) {
	query := "SELECT * FROM day WHERE id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, name)

	day := models.Day{}
	err := row.Scan(
		&day.Id,
		&day.Name,
		&day.Active,
		&day.CreatedAt,
		&day.UpdatedAt)

	if err != nil {
		return day, err
	}

	return day, nil
}

func (r *DayRepository) GetById(ctx context.Context, id string) (models.Day, error) {
	query := "SELECT id, name FROM day WHERE id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	day := models.Day{}
	err := row.Scan(&day.Id, &day.Name)

	if err != nil {
		return day, err
	}

	return day, nil
}

func (r *DayRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM day;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
