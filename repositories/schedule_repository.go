package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
	"time"
)

type ScheduleRepository struct {
	Data *db.Data
}

func (r *ScheduleRepository) GetAll(ctx context.Context, offset, limit int) ([]models.Schedule, error) {
	query := "SELECT * FROM schedule OFFSET($1) LIMIT($2);"
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var schedules []models.Schedule

	for rows.Next() {
		var schedule models.Schedule

		rows.Scan(
			&schedule.Id,
			&schedule.From,
			&schedule.To,
			&schedule.ProfileId,
			&schedule.DayId,
			&schedule.CreatedAt,
			&schedule.UpdatedAt)

		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetOne(ctx context.Context, id string) (models.Schedule, error) {
	query := "SELECT * FROM schedules WHERE id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var schedule models.Schedule
	err := row.Scan(
		&schedule.Id,
		&schedule.From,
		&schedule.To,
		&schedule.ProfileId,
		&schedule.DayId,
		&schedule.CreatedAt,
		&schedule.UpdatedAt)

	if err != nil {
		return models.Schedule{}, err
	}

	return schedule, nil
}

func (r *ScheduleRepository) Create(ctx context.Context, schedule *models.Schedule) error {
	query := "INSERT INTO schedules VALUES ($1, $2, $3, $4, $5, $6, $7);"
	schedule.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		schedule.Id,
		schedule.From,
		schedule.To,
		schedule.ProfileId,
		schedule.DayId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *ScheduleRepository) DeleteOne(ctx context.Context, id string) error {
	query := "DELETE FROM schedules WHERE id = $1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *ScheduleRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM schedules;"
	row := r.Data.DB.QueryRowContext(ctx, query)

	var number int

	row.Scan(&number)

	return number, nil
}
