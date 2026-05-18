package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	models "github.com/Raphsodyz/spacedevs-go/models"
	"github.com/redis/go-redis/v9"
)

const (
	searchCacheTTL    = 5 * time.Minute
	searchCachePrefix = "search:launch:"
)

type RedisRepository interface {
	GetFromSearch(ctx context.Context, search models.SearchLaunchRequest) (*models.SearchLaunchResponse, error)
	SetSearchPagination(ctx context.Context, search models.SearchLaunchRequest, result *models.SearchLaunchResponse) error
}

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) RedisRepository {
	return &redisRepository{client: client}
}

func (r *redisRepository) GetFromSearch(ctx context.Context, search models.SearchLaunchRequest) (*models.SearchLaunchResponse, error) {
	key, err := buildSearchCacheKey(search)
	if err != nil {
		return nil, err
	}

	val, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var result models.SearchLaunchResponse
	if err := json.Unmarshal(val, &result); err != nil {
		return nil, fmt.Errorf("redisRepository.GetFromSearch unmarshal: %w", err)
	}

	return &result, nil
}

func (r *redisRepository) SetSearchPagination(ctx context.Context, search models.SearchLaunchRequest, result *models.SearchLaunchResponse) error {
	key, err := buildSearchCacheKey(search)
	if err != nil {
		return err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("redisRepository.SetSearchPagination marshal: %w", err)
	}

	if err := r.client.Set(ctx, key, data, searchCacheTTL).Err(); err != nil {
		return fmt.Errorf("redisRepository.SetSearchPagination set: %w", err)
	}

	return nil
}

func buildSearchCacheKey(search models.SearchLaunchRequest) (string, error) {
	if search.Mission == nil &&
		search.Rocket == nil &&
		search.Location == nil &&
		search.Pad == nil &&
		search.Launch == nil {
		return "", fmt.Errorf("redisRepository.buildSearchCacheKey: no search filters provided")
	}

	var sb strings.Builder

	if search.Mission != nil && strings.TrimSpace(*search.Mission) != "" {
		sb.WriteString("mission:" + strings.TrimSpace(*search.Mission) + "_")
	}
	if search.Rocket != nil && strings.TrimSpace(*search.Rocket) != "" {
		sb.WriteString("rocket:" + strings.TrimSpace(*search.Rocket) + "_")
	}
	if search.Location != nil && strings.TrimSpace(*search.Location) != "" {
		sb.WriteString("location:" + strings.TrimSpace(*search.Location) + "_")
	}
	if search.Pad != nil && strings.TrimSpace(*search.Pad) != "" {
		sb.WriteString("pad:" + strings.TrimSpace(*search.Pad) + "_")
	}
	if search.Launch != nil && strings.TrimSpace(*search.Launch) != "" {
		sb.WriteString("launch:" + strings.TrimSpace(*search.Launch) + "_")
	}

	page := 0
	if search.Page != nil {
		page = *search.Page
	}

	return fmt.Sprintf("%s%s_page:%d", searchCachePrefix, strings.TrimRight(sb.String(), "_"), page), nil
}
