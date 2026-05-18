package usecase

import (
	"context"
	"fmt"
	"strings"
	"sync"

	models "github.com/Raphsodyz/spacedevs-go/models"
	repository "github.com/Raphsodyz/spacedevs-go/repository"
)

type SearchLaunchUseCase struct {
	missionRepo       repository.MissionRepository
	configurationRepo repository.ConfigurationRepository
	locationRepo      repository.LocationRepository
	padRepo           repository.PadRepository
	launchRepo        repository.LaunchRepository
	redisRepo         repository.RedisRepository
}

func NewSearchLaunchUseCase(
	missionRepo repository.MissionRepository,
	configurationRepo repository.ConfigurationRepository,
	locationRepo repository.LocationRepository,
	padRepo repository.PadRepository,
	launchRepo repository.LaunchRepository,
	redisRepo repository.RedisRepository,
) *SearchLaunchUseCase {
	return &SearchLaunchUseCase{
		missionRepo:       missionRepo,
		configurationRepo: configurationRepo,
		locationRepo:      locationRepo,
		padRepo:           padRepo,
		launchRepo:        launchRepo,
		redisRepo:         redisRepo,
	}
}

func (uc *SearchLaunchUseCase) SearchByRequest(search models.SearchLaunchRequest) (*models.SearchLaunchResponse, error) {
	ctx := context.Background()

	cached, err := uc.redisRepo.GetFromSearch(ctx, search)
	if err == nil && cached != nil {
		return cached, nil
	}

	filters := &models.SearchLaunchFilters{}

	var wg sync.WaitGroup
	var mu sync.Mutex
	errCh := make(chan error, 5)

	if search.Mission != nil && strings.TrimSpace(*search.Mission) != "" {
		wg.Add(1)
		go func(mission string) {
			defer wg.Done()

			missionIDs, err := uc.missionRepo.GetIdsByMissionName(ctx, mission)
			if err != nil {
				errCh <- fmt.Errorf("failed to get mission IDs: %w", err)
				return
			}

			mu.Lock()
			filters.MissionIds = missionIDs
			mu.Unlock()
		}(*search.Mission)
	}

	if search.Rocket != nil && strings.TrimSpace(*search.Rocket) != "" {
		wg.Add(1)
		go func(rocket string) {
			defer wg.Done()

			rocketIDs, err := uc.configurationRepo.GetIdsByRocketName(ctx, rocket)
			if err != nil {
				errCh <- fmt.Errorf("failed to get configuration IDs: %w", err)
				return
			}

			mu.Lock()
			filters.RocketIds = rocketIDs
			mu.Unlock()
		}(*search.Rocket)
	}

	if search.Location != nil && strings.TrimSpace(*search.Location) != "" {
		wg.Add(1)
		go func(location string) {
			defer wg.Done()

			locationIDs, err := uc.locationRepo.GetIdsByLocationName(ctx, location)
			if err != nil {
				errCh <- fmt.Errorf("failed to get location IDs: %w", err)
				return
			}

			mu.Lock()
			filters.LocationIds = locationIDs
			mu.Unlock()
		}(*search.Location)
	}

	if search.Pad != nil && strings.TrimSpace(*search.Pad) != "" {
		wg.Add(1)
		go func(pad string) {
			defer wg.Done()

			padIDs, err := uc.padRepo.GetIdsByPadName(ctx, pad)
			if err != nil {
				errCh <- fmt.Errorf("failed to get pad IDs: %w", err)
				return
			}

			mu.Lock()
			filters.PadIds = padIDs
			mu.Unlock()
		}(*search.Pad)
	}

	if search.Launch != nil && strings.TrimSpace(*search.Launch) != "" {
		wg.Add(1)
		go func(launch string) {
			defer wg.Done()

			launchIDs, err := uc.launchRepo.GetIdsBySlugName(ctx, launch)
			if err != nil {
				errCh <- fmt.Errorf("failed to get launch IDs: %w", err)
				return
			}

			mu.Lock()
			filters.LaunchIds = launchIDs
			mu.Unlock()
		}(*search.Launch)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	result, err := uc.launchRepo.GetSearchResults(ctx, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to get search results: %w", err)
	}

	if setErr := uc.redisRepo.SetSearchPagination(ctx, search, result); setErr != nil {
		fmt.Printf("search_launch_usecase.SearchByRequest: cache write failed: %v\n", setErr)
	}

	return result, nil
}
