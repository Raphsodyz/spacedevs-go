package entity

import (
	audit "github.com/spacedevs-go/internal/domain/entity/Audit"
)

type Mission struct {
	Id               int64   `json:"id" db:"id" validate:"omitempty,int64"`
	IdFromApi        *int64  `json:"id_from_api" db:"id_from_api" validate:"omitempty,int64"`
	LaunchLibraryId  *int64  `json:"launch_library_id" db:"launch_library_id" validate:"omitempty,int64"`
	Name             *string `json:"name" db:"name" validate:"omitempty,lte=360"`
	Description      *string `json:"description" db:"description" validate:"omitempty,lte=5000"`
	Type             *string `json:"type" db:"type" validate:"omitempty,lte=360"`
	IdOrbit          *int64  `json:"id_orbit" db:"id_orbit" validate:"omitempty,int64"`
	Orbit            *Orbit  `json:"orbit,omitempty" db:"-"`
	LaunchDesignator *string `json:"launch_designator" db:"launch_designator" validate:"omitempty,lte=360"`
	Search           *string `json:"-" db:"search" validate:"omitempty,lte=360"`
	audit.Audit
}
