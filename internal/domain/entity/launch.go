package entity

import (
	"time"

	"github.com/google/uuid"
	audit "github.com/spacedevs-go/internal/domain/entity/Audit"
)

type Launch struct {
	Id                      int64                  `json:"id" db:"id" validate:"omitempty,int64"`
	IdFromApi               *int64                 `json:"id_from_api" db:"id_from_api" validate:"omitempty,int64"`
	ApiUuid                 *uuid.UUID             `json:"api_uuid" db:"api_uuid" validate:"omitempty,uuid"`
	Url                     *string                `json:"url" db:"url" validate:"omitempty,lte=1000"`
	LaunchLibraryId         *int64                 `json:"launch_library_id" db:"launch_library_id" validate:"omitempty,int64"`
	Slug                    *string                `json:"slug" db:"slug" validate:"omitempty,lte=360"`
	Name                    *string                `json:"name" db:"name" validate:"omitempty,lte=360"`
	IdStatus                *int64                 `json:"id_status" db:"id_status" validate:"omitempty,int64"`
	Status                  *Status                `json:"status,omitempty" db:"-"`
	Net                     *time.Time             `json:"net" db:"net" validate:"omitempty"`
	WindowEnd               *time.Time             `json:"window_end" db:"window_end" validate:"omitempty"`
	WindowStart             *time.Time             `json:"window_start" db:"window_start" validate:"omitempty"`
	Inhold                  *bool                  `json:"inhold" db:"inhold" validate:"omitempty"`
	TbdTime                 *bool                  `json:"tbd_time" db:"tbd_time" validate:"omitempty"`
	TbdDate                 *bool                  `json:"tbd_date" db:"tbd_date" validate:"omitempty"`
	Probability             *int64                 `json:"probability" db:"probability" validate:"omitempty,int64"`
	HoldReason              *string                `json:"hold_reason" db:"hold_reason" validate:"omitempty,lte=360"`
	FailReason              *string                `json:"fail_reason" db:"fail_reason" validate:"omitempty,lte=750"`
	Hashtag                 *string                `json:"hashtag" db:"hashtag" validate:"omitempty,lte=360"`
	IdLaunchServiceProvider *int64                 `json:"id_launch_service_provider" db:"id_launch_service_provider" validate:"omitempty,int64"`
	LaunchServiceProvider   *LaunchServiceProvider `json:"launch_service_provider,omitempty" db:"-"`
	IdRocket                *int64                 `json:"id_rocket" db:"id_rocket" validate:"omitempty,int64"`
	Rocket                  *Rocket                `json:"rocket,omitempty" db:"-"`
	IdMission               *int64                 `json:"id_mission" db:"id_mission" validate:"omitempty,int64"`
	Mission                 *Mission               `json:"mission,omitempty" db:"-"`
	IdPad                   *int64                 `json:"id_pad" db:"id_pad" validate:"omitempty,int64"`
	Pad                     *Pad                   `json:"pad,omitempty" db:"-"`
	WebcastLive             *bool                  `json:"webcast_live" db:"web_cast_live" validate:"omitempty"`
	Image                   *string                `json:"image" db:"image" validate:"omitempty,lte=1000"`
	Infographic             *string                `json:"infographic" db:"infographic" validate:"omitempty,lte=360"`
	Programs                *string                `json:"programs" db:"programs" validate:"omitempty,lte=360"`
	Search                  *string                `json:"-" db:"search" validate:"omitempty,lte=600"`
	audit.Audit
}
