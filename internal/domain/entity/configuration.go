package entity

import (
	audit "github.com/spacedevs-go/internal/domain/entity/Audit"
)

type Configuration struct {
	Id              int64   `json:"id" db:"id" validate:"omitempty,int64"`
	IdFromApi       *int64  `json:"id_from_api" db:"id_from_api" validate:"omitempty,int64"`
	LaunchLibraryId *int64  `json:"launch_library_id" db:"launch_library_id" validate:"omitempty,int64"`
	Url             *string `json:"url" db:"url" validate:"omitempty,lte=1000"`
	Name            *string `json:"name" db:"name" validate:"omitempty,lte=360"`
	Family          *string `json:"family" db:"family" validate:"omitempty,lte=360"`
	FullName        *string `json:"full_name" db:"full_name" validate:"omitempty,lte=360"`
	Variant         *string `json:"variant" db:"variant" validate:"omitempty,lte=360"`
	Search          *string `json:"-" db:"search" validate:"omitempty,lte=600"`
	audit.Audit
}
