package entity

import (
	ProcessingStatus "github.com/Raphsodyz/spacedevs-go/internal/enum"
	Audit "github.com/Raphsodyz/spacedevs-go/models"
)

type Orbit struct {
	Id        int64   `json:"id" db:"id" validate:"omitempty,int64"`
	IdFromApi *int64  `json:"id_from_api" db:"id_from_api" validate:"omitempty,int64"`
	Name      *string `json:"name" db:"name" validate:"omitempty,lte=360"`
	Abbrev    *string `json:"abbrev" db:"abbrev" validate:"omitempty,lte=360"`
	Audit.Audit
	ProcessingStatus.ProcessingStatus
}
