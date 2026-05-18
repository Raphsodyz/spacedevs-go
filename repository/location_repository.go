package repository

import (
	"context"
	"fmt"
	"time"

	squirrel "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LocationRepository interface {
	GetIdsByLocationName(ctx context.Context, location string) ([]int64, error)
}

type locationRepository struct {
	db *pgxpool.Pool
}

func NewLocationRepository(db *pgxpool.Pool) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) GetIdsByLocationName(ctx context.Context, location string) ([]int64, error) {
	var pgsql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := pgsql.Select(
		"l.id",
	).
		From("data.location AS l").
		Where(squirrel.ILike{"l.search": location}).
		Where(squirrel.Eq{"l.status": "PUBLISHED"}).
		Where(squirrel.Eq{"l.effective_date": time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("locationRepository.GetIdsByLocationName data build: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("locationRepository.GetIdsByLocationName data exec: %w", err)
	}

	defer rows.Close()

	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, fmt.Errorf("locationRepository.GetIdsByLocationName collect: %w", err)
	}

	return ids, nil
}
