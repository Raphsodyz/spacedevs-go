package usecase

import (
	repository "github.com/Raphsodyz/spacedevs-go/repository"
)

type SearchLaunchUseCase struct {
	missionRepo       repository.MissionRepository
	configurationRepo repository.ConfigurationRepository
	locationRepo      repository.LocationRepository
	padRepo           repository.PadRepository
	launchRepo        repository.LaunchRepository
}

func NewSearchLaunchUseCase(
	missionRepo repository.MissionRepository,
	configurationRepo repository.ConfigurationRepository,
	locationRepo repository.LocationRepository,
	padRepo repository.PadRepository,
	launchRepo repository.LaunchRepository,
) *SearchLaunchUseCase {
	return &SearchLaunchUseCase{
		missionRepo:       missionRepo,
		configurationRepo: configurationRepo,
		locationRepo:      locationRepo,
		padRepo:           padRepo,
		launchRepo:        launchRepo,
	}
}
