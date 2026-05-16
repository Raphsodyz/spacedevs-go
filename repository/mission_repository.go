package repository

import (
	"context"
	"fmt"
	"time"

	squirrel "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionRepository interface {
	GetIdsByMissionName(ctx context.Context, mission string) ([]int64, error)
}

type missionRepository struct {
	db *pgxpool.Pool
}

func NewMissionRepository(db *pgxpool.Pool) MissionRepository {
	return &missionRepository{db: db}
}

func (r *missionRepository) GetIdsByMissionName(ctx context.Context, mission string) ([]int64, error) {
	var pgsql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := pgsql.Select(
		"m.id",
	).
		From("public.mission AS m").
		Where(squirrel.ILike{"m.search": mission}).
		Where(squirrel.Eq{"m.status": "PUBLISHED"}).
		Where(squirrel.Eq{"m.effective_date": time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("missionRepository.GetIdsByMissionName data build: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("missionRepository.GetIdsByMissionName data exec: %w", err)
	}

	defer rows.Close()

	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, fmt.Errorf("missionRepository.GetIdsByMissionName collect: %w", err)
	}

	return ids, nil
}
