package entity

import (
	ProcessingStatus "github.com/spacedevs-go/internal/enum"
	Audit "github.com/spacedevs-go/model"
)

type Status struct {
	Id          int64   `json:"id" db:"id" validate:"omitempty,int64"`
	IdFromApi   *int64  `json:"id_from_api" db:"id_from_api" validate:"omitempty,int64"`
	Name        *string `json:"name" db:"name" validate:"omitempty,lte=360"`
	Abbrev      *string `json:"abbrev" db:"abbrev" validate:"omitempty,lte=360"`
	Description *string `json:"description" db:"description" validate:"omitempty,lte=5000"`
	Audit.Audit
	ProcessingStatus.ProcessingStatus
}
