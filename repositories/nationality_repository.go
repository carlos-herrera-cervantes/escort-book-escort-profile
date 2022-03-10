package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
)

type NationalityRepository struct {
	Data *db.Data
}

func (r *NationalityRepository) GetAll(ctx context.Context, offset, limit int) ([]models.Nationality, error) {
	query := "SELECT * FROM nationality OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var nationalities []models.Nationality

	for rows.Next() {
		var nationality models.Nationality

		rows.Scan(
			&nationality.Id,
			&nationality.Name,
			&nationality.Active,
			&nationality.CreatedAt,
			&nationality.UpdatedAt)

		nationalities = append(nationalities, nationality)
	}

	return nationalities, nil
}

func (r *NationalityRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM nationality;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
