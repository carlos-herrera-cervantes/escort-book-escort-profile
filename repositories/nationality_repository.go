package repositories

import (
	"context"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type NationalityRepository struct {
	Data *singleton.PostgresClient
}

func (r *NationalityRepository) GetAll(ctx context.Context, offset, limit int) ([]models.Nationality, error) {
	query := "SELECT * FROM nationality OFFSET($1) LIMIT($2);"
	rows, err := r.Data.EscortProfileDB.QueryContext(ctx, query, offset, limit)
	nationalities := []models.Nationality{}

	if err != nil {
		return nationalities, err
	}

	defer rows.Close()

	for rows.Next() {
		nationality := models.Nationality{}

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

func (r *NationalityRepository) GetById(ctx context.Context, id string) (models.Nationality, error) {
	query := "SELECT id, name from nationality WHERE id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	nationality := models.Nationality{}
	err := row.Scan(&nationality.Id, &nationality.Name)

	if err != nil {
		return nationality, err
	}

	return nationality, nil
}

func (r *NationalityRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM nationality;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
