package entity

import (
	"github.com/google/uuid"
	audit "github.com/spacedevs-go/internal/domain/entity/Audit"
)

type Launch struct {
	Id              int64      `json:"id" db:"id"  validate:"omitempty,int64"`
	ApiUuid         *uuid.UUID `json:"api_uuid" db:"api_uuid" validate:"omitempy,uuid"`
	Url             *string    `json:"url" db:"url" validate:"lte=1000"`
	LaunchLibraryId *int64     `json:"launch_library_id" db:"launch_library_id" validate:"omitempy,int64"`
	audit.Audit
}
