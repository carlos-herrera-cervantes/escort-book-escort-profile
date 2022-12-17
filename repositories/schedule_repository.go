package repositories

import (
	"context"
	"time"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"
)

type ScheduleRepository struct {
	Data *singleton.PostgresClient
}

func (r *ScheduleRepository) GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.Schedule, error) {
	query := `SELECT a.id, a.from, a.to, a.escort_id, a.day_id, a.created_at, a.updated_at, b.name
			  FROM schedule a
			  INNER JOIN day b
			  ON b.id = a.day_id
			  WHERE escort_id = $3 OFFSET($1) LIMIT($2);`
	rows, err := r.Data.EscortProfileDB.QueryContext(ctx, query, offset, limit, profileId)
	schedules := []models.Schedule{}

	if err != nil {
		return schedules, err
	}

	defer rows.Close()

	for rows.Next() {
		schedule := models.Schedule{}

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
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, id)

	schedule := models.Schedule{}
	err := row.Scan(
		&schedule.Id,
		&schedule.From,
		&schedule.To,
		&schedule.ProfileId,
		&schedule.DayId,
		&schedule.CreatedAt,
		&schedule.UpdatedAt)

	if err != nil {
		return schedule, err
	}

	return schedule, nil
}

func (r *ScheduleRepository) Create(ctx context.Context, schedule *models.Schedule) error {
	query := "INSERT INTO schedule VALUES ($1, $2, $3, $4, $5, $6, $7);"
	schedule.SetDefaultValues()

	_, err := r.Data.EscortProfileDB.ExecContext(
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
	_, err := r.Data.EscortProfileDB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *ScheduleRepository) Count(ctx context.Context, profileId string) (int, error) {
	query := "SELECT COUNT(*) FROM schedule WHERE escort_id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, profileId)

	var number int

	row.Scan(&number)

	return number, nil
}
