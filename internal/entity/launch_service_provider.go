package entity

import (
	ProcessingStatus "github.com/Raphsodyz/spacedevs-go/internal/enum"
	audit "github.com/Raphsodyz/spacedevs-go/models"
)

type LaunchServiceProvider struct {
	Id        int64   `json:"id" db:"id" validate:"omitempty,int64"`
	IdFromApi *int64  `json:"id_from_api" db:"id_from_api" validate:"omitempty,int64"`
	Url       *string `json:"url" db:"url" validate:"omitempty,lte=1000"`
	Name      *string `json:"name" db:"name" validate:"omitempty,lte=360"`
	Type      *string `json:"type" db:"type" validate:"omitempty,lte=360"`
	audit.Audit
	ProcessingStatus.ProcessingStatus
}
