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

func (r *ScheduleRepository) GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.Schedule, error) {
	query := `SELECT a.id, a.from, a.to, a.escort_id, a.day_id, a.created_at, a.updated_at, b.name
			  FROM schedule a
			  INNER JOIN day b
			  ON b.id = a.day_id
			  WHERE escort_id = $3 OFFSET($1) LIMIT($2);`
	rows, err := r.Data.DB.QueryContext(ctx, query, offset, limit, profileId)

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
			&schedule.UpdatedAt,
			&schedule.DayName,
		)

		schedule.From = schedule.From[11:19]
		schedule.To = schedule.To[11:19]

		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (r *ScheduleRepository) GetOne(ctx context.Context, id string) (models.Schedule, error) {
	query := "SELECT * FROM schedule WHERE id = $1;"
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
	query := "INSERT INTO schedule VALUES ($1, $2, $3, $4, $5, $6, $7);"
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
	query := "DELETE FROM schedule WHERE id = $1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *ScheduleRepository) Count(ctx context.Context, profileId string) (int, error) {
	query := "SELECT COUNT(*) FROM schedule WHERE escort_id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, profileId)

	var number int

	row.Scan(&number)

	return number, nil
}
