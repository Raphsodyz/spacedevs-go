package entity

import (
	Audit "github.com/spacedevs-go/model"
)

type Location struct {
	Id                int64   `json:"id" db:"id" validate:"omitempty,int64"`
	IdFromApi         *int64  `json:"id_from_api" db:"id_from_api" validate:"omitempty,int64"`
	Url               *string `json:"url" db:"url" validate:"omitempty,lte=1000"`
	Name              *string `json:"name" db:"name" validate:"omitempty,lte=360"`
	CountryCode       *string `json:"country_code" db:"country_code" validate:"omitempty,lte=360"`
	MapImage          *string `json:"map_image" db:"map_image" validate:"omitempty,lte=1000"`
	TotalLaunchCount  *int64  `json:"total_launch_count" db:"total_launch_count" validate:"omitempty,int64"`
	TotalLandingCount *int64  `json:"total_landing_count" db:"total_landing_count" validate:"omitempty,int64"`
	Search            *string `json:"-" db:"search" validate:"omitempty,lte=360"`
	Audit.Audit
}
