package entity

import (
	ProcessingStatus "github.com/spacedevs-go/internal/enum"
	Audit "github.com/spacedevs-go/model"
)

type Pad struct {
	Id               int64     `json:"id" db:"id" validate:"omitempty,int64"`
	IdFromApi        *int64    `json:"id_from_api" db:"id_from_api" validate:"omitempty,int64"`
	Url              *string   `json:"url" db:"url" validate:"omitempty,lte=1000"`
	AgencyId         *int64    `json:"agency_id" db:"agency_id" validate:"omitempty,int64"`
	Name             *string   `json:"name" db:"name" validate:"omitempty,lte=360"`
	InfoUrl          *string   `json:"info_url" db:"info_url" validate:"omitempty,lte=1000"`
	WikiUrl          *string   `json:"wiki_url" db:"wiki_url" validate:"omitempty,lte=1000"`
	MapUrl           *string   `json:"map_url" db:"map_url" validate:"omitempty,lte=1000"`
	Latitude         *float64  `json:"latitude" db:"latitude" validate:"omitempty"`
	Longitude        *float64  `json:"longitude" db:"longitude" validate:"omitempty"`
	IdLocation       *int64    `json:"id_location" db:"id_location" validate:"omitempty,int64"`
	Location         *Location `json:"location,omitempty" db:"-"`
	MapImage         *string   `json:"map_image" db:"map_image" validate:"omitempty,lte=1000"`
	TotalLaunchCount *int64    `json:"total_launch_count" db:"total_launch_count" validate:"omitempty,int64"`
	Search           *string   `json:"-" db:"search" validate:"omitempty,lte=360"`
	Audit.Audit
	ProcessingStatus.ProcessingStatus
}
