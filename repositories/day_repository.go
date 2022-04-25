package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
)

type DayRepository struct {
	Data *db.Data
}

func (r *DayRepository) GetAll(ctx context.Context, offset, limit int) ([]models.Day, error) {
	query := "SELECT * FROM day OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var days []models.Day

	for rows.Next() {
		var day models.Day

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
	row := r.Data.DB.QueryRowContext(ctx, query, name)

	var day models.Day
	err := row.Scan(
		&day.Id,
		&day.Name,
		&day.Active,
		&day.CreatedAt,
		&day.UpdatedAt)

	if err != nil {
		return models.Day{}, err
	}

	return day, nil
}

func (r *DayRepository) GetById(ctx context.Context, id string) (models.Day, error) {
	query := "SELECT id, name FROM day WHERE id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var day models.Day
	err := row.Scan(&day.Id, &day.Name)

	if err != nil {
		return models.Day{}, err
	}

	return day, nil
}

func (r *DayRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM day;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
