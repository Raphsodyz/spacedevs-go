package repository

import (
	"context"
	"fmt"
	"time"

	squirrel "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConfigurationRepository interface {
	GetIdsByRocketName(ctx context.Context, rocket string) ([]int64, error)
}

type configurationRepository struct {
	db *pgxpool.Pool
}

func NewConfigurationRepository(db *pgxpool.Pool) ConfigurationRepository {
	return &configurationRepository{db: db}
}

func (r *configurationRepository) GetIdsByRocketName(ctx context.Context, rocket string) ([]int64, error) {
	var pgsql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := pgsql.Select(
		"c.id",
	).
		From("public.configuration AS c").
		Join("public.rocket AS r ON c.id = r.id_configuration").
		Where(squirrel.ILike{"r.search": rocket}).
		Where(squirrel.Eq{"c.status": "PUBLISHED"}).
		Where(squirrel.Eq{"c.effective_date": time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("configurationRepository.GetIdsByRocketName data build: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("configurationRepository.GetIdsByRocketName data exec: %w", err)
	}

	defer rows.Close()

	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, fmt.Errorf("configurationRepository.GetIdsByRocketName collect: %w", err)
	}

	return ids, nil
}
