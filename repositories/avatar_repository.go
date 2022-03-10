package repositories

import (
	"context"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/models"
	"time"
)

type AvatarRepository struct {
	Data *db.Data
}

func (r *AvatarRepository) GetOne(ctx context.Context, id string) (models.Avatar, error) {
	query := "SELECT * FROM avatar WHERE escort_id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, id)

	var avatar models.Avatar
	err := row.Scan(
		&avatar.Id,
		&avatar.Path,
		&avatar.ProfileId,
		&avatar.CreatedAt,
		&avatar.UpdatedAt)

	if err != nil {
		return models.Avatar{}, err
	}

	return avatar, nil
}

func (r *AvatarRepository) Create(ctx context.Context, avatar *models.Avatar) error {
	query := "INSERT INTO avatar VALUES ($1, $2, $3, $4, $5);"
	avatar.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		avatar.Id,
		avatar.Path,
		avatar.ProfileId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *AvatarRepository) UpdateOne(ctx context.Context, id string, avatar *models.Avatar) error {
	query := "UPDATE avatar SET path = $1, updated_at = $2 WHERE escort_id = $3;"
	_, err := r.Data.DB.ExecContext(ctx, query, avatar.Path, time.Now().UTC(), id)

	if err != nil {
		return err
	}

	return nil
}

func (r *AvatarRepository) DeleteOne(ctx context.Context, id string) error {
	query := "DELETE FROM avatar WHERE escort_id = $1;"
	_, err := r.Data.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
