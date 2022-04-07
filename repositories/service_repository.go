package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
	"time"
)

type ServiceRepository struct {
	Data *db.Data
}

func (r *ServiceRepository) GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.Service, error) {
	query := `SELECT a.*, b.name FROM service a
				INNER JOIN service_category b
				ON a.service_category_id = b.id
				WHERE escort_id = $3 OFFSET($1) LIMIT($2);`
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit, profileId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var services []models.Service

	for rows.Next() {
		var service models.Service

		rows.Scan(
			&service.Id,
			&service.ProfileId,
			&service.ServiceCategoryId,
			&service.CreatedAt,
			&service.UpdatedAt,
			&service.Cost,
			&service.Name,
		)

		services = append(services, service)
	}

	return services, nil
}

func (r *ServiceRepository) GetOne(ctx context.Context, id, profileId string) (models.Service, error) {
	query := "SELECT * FROM service WHERE id = $1 AND escort_id = $2;"
	row := r.Data.DB.QueryRowContext(ctx, query, id, profileId)

	var service models.Service
	err := row.Scan(
		&service.Id,
		&service.ProfileId,
		&service.ServiceCategoryId,
		&service.CreatedAt,
		&service.UpdatedAt,
		&service.Cost,
	)

	if err != nil {
		return models.Service{}, err
	}

	return service, nil
}

func (r *ServiceRepository) Create(ctx context.Context, service *models.Service) error {
	query := "INSERT INTO service VALUES ($1, $2, $3, $4, $5, $6);"
	service.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		service.Id,
		service.ProfileId,
		service.ServiceCategoryId,
		time.Now().UTC(),
		time.Now().UTC(),
		service.Cost,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ServiceRepository) DeleteOne(ctx context.Context, id, profileId string) error {
	query := "DELETE FROM service WHERE id = $1 AND escort_id = $2;"
	_, err := r.Data.DB.ExecContext(ctx, query, id, profileId)

	if err != nil {
		return err
	}

	return nil
}

func (r *ServiceRepository) Count(ctx context.Context, profileId string) (int, error) {
	query := "SELECT COUNT(*) FROM service WHERE escort_id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, profileId)

	var number int

	row.Scan(&number)

	return number, nil
}
