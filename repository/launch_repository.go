package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	entity "github.com/spacedevs-go/internal/entity"
)

type LaunchRepository interface {
	GetById(ctx context.Context, launchId int64) (*entity.Launch, error)
}

type launchRepository struct {
	db *pgxpool.Pool
}

func (r *launchRepository) GetById(ctx context.Context, launchId int64) (*entity.Launch, error) {
	query := ``SELECT * FROM `
}
