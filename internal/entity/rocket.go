package entity

import (
	ProcessingStatus "github.com/Raphsodyz/spacedevs-go/internal/enum"
	Audit "github.com/Raphsodyz/spacedevs-go/models"
)

type Rocket struct {
	Id              int64          `json:"id" db:"id" validate:"omitempty,int64"`
	IdFromApi       *int64         `json:"id_from_api" db:"id_from_api" validate:"omitempty,int64"`
	IdConfiguration *int64         `json:"id_configuration" db:"id_configuration" validate:"omitempty,int64"`
	Configuration   *Configuration `json:"configuration,omitempty" db:"-"`
	Audit.Audit
	ProcessingStatus.ProcessingStatus
}
