package repository

import (
	"context"
	"fmt"
	"time"

	squirrel "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PadRepository interface {
	GetIdsByPadName(ctx context.Context, padName string) ([]int64, error)
}

type padRepository struct {
	db *pgxpool.Pool
}

func (r *padRepository) GetIdsByPadName(ctx context.Context, padName string) ([]int64, error) {
	var pgsql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := pgsql.Select(
		"p.id",
	).
		From("public.pad AS p").
		Where(squirrel.ILike{"p.name": padName}).
		Where(squirrel.Eq{"p.status": "PUBLISHED"}).
		Where(squirrel.Eq{"p.effective_date": time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("padRepository.GetIdsByPadName data build: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("padRepository.GetIdsByPadName data exec: %w", err)
	}

	defer rows.Close()

	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, fmt.Errorf("padRepository.GetIdsByPadName collect: %w", err)
	}

	return ids, nil
}
