package models

import (
	"time"

	"github.com/google/uuid"
)

type LaunchView struct {
	Id                           uuid.UUID  `json:"launch_id" db:"launch_id"`
	AtualizationDate             time.Time  `json:"launch_atualization_date" db:"launch_atualization_date"`
	ImportedT                    time.Time  `json:"launch_imported_t" db:"launch_imported_t"`
	EntityStatus                 string     `json:"launch_status" db:"launch_status"`
	ApiGuid                      uuid.UUID  `json:"launch_api_guid" db:"launch_api_guid"`
	Url                          *string    `json:"launch_url" db:"launch_url"`
	LaunchLibraryId              *int64     `json:"launch_launch_library_id" db:"launch_launch_library_id"`
	Slug                         *string    `json:"launch_slug" db:"launch_slug"`
	Name                         *string    `json:"launch_name" db:"launch_name"`
	IdStatus                     *uuid.UUID `json:"id_status" db:"id_status"`
	Net                          *time.Time `json:"launch_net" db:"launch_net"`
	WindowEnd                    *time.Time `json:"launch_window_end" db:"launch_window_end"`
	WindowStart                  *time.Time `json:"launch_window_start" db:"launch_window_start"`
	Inhold                       *bool      `json:"launch_inhold" db:"launch_inhold"`
	TbdTime                      *bool      `json:"launch_tbd_time" db:"launch_tbd_time"`
	TbdDate                      *bool      `json:"launch_tbd_date" db:"launch_tbd_date"`
	Probability                  *int32     `json:"launch_probability" db:"launch_probability"`
	HoldReason                   *string    `json:"launch_hold_reason" db:"launch_hold_reason"`
	FailReason                   *string    `json:"launch_fail_reason" db:"launch_fail_reason"`
	Hashtag                      *string    `json:"launch_hashtag" db:"launch_hashtag"`
	IdLaunchServiceProvider      *uuid.UUID `json:"id_launch_service_provider" db:"id_launch_service_provider"`
	IdRocket                     *uuid.UUID `json:"id_rocket" db:"id_rocket"`
	IdMission                    *uuid.UUID `json:"id_mission" db:"id_mission"`
	IdPad                        *uuid.UUID `json:"id_pad" db:"id_pad"`
	WebcastLive                  *bool      `json:"launch_web_cast_live" db:"launch_web_cast_live"`
	Image                        *string    `json:"launch_image" db:"launch_image"`
	Infographic                  *string    `json:"launch_infographic" db:"launch_infographic"`
	Programs                     *string    `json:"launch_programs" db:"launch_programs"`
	StatusName                   *string    `json:"status_name" db:"status_name"`
	StatusAbbrev                 *string    `json:"status_abbrev" db:"status_abbrev"`
	StatusDescription            *string    `json:"status_description" db:"status_description"`
	LaunchServiceProviderUrl     *string    `json:"launch_service_provider_url" db:"launch_service_provider_url"`
	LaunchServiceProviderName    *string    `json:"launch_service_provider_name" db:"launch_service_provider_name"`
	LaunchServiceProviderType    *string    `json:"launch_service_provider_type" db:"launch_service_provider_type"`
	IdConfiguration              *uuid.UUID `json:"id_configuration" db:"id_configuration"`
	ConfigurationLaunchLibraryId *int64     `json:"configuration_launch_library_id" db:"configuration_launch_library_id"`
	ConfigurationUrl             *string    `json:"configuration_url" db:"configuration_url"`
	ConfigurationName            *string    `json:"configuration_name" db:"configuration_name"`
	ConfigurationFamily          *string    `json:"configuration_family" db:"configuration_family"`
	ConfigurationFullName        *string    `json:"configuration_full_name" db:"configuration_full_name"`
	ConfigurationVariant         *string    `json:"configuration_variant" db:"configuration_variant"`
	MissionLaunchLibraryId       *int64     `json:"mission_launch_library_id" db:"mission_launch_library_id"`
	MissionName                  *string    `json:"mission_name" db:"mission_name"`
	MissionDescription           *string    `json:"mission_description" db:"mission_description"`
	MissionType                  *string    `json:"mission_type" db:"mission_type"`
	IdOrbit                      *uuid.UUID `json:"id_orbit" db:"id_orbit"`
	MissionLaunchDesignator      *string    `json:"mission_launch_designator" db:"mission_launch_designator"`
	OrbitName                    *string    `json:"orbit_name" db:"orbit_name"`
	OrbitAbbrev                  *string    `json:"orbit_abbrev" db:"orbit_abbrev"`
	PadUrl                       *string    `json:"pad_url" db:"pad_url"`
	PadAgencyId                  *int32     `json:"pad_agency_id" db:"pad_agency_id"`
	PadName                      *string    `json:"pad_name" db:"pad_name"`
	PadInfoUrl                   *string    `json:"pad_info_url" db:"pad_info_url"`
	PadWikiUrl                   *string    `json:"pad_wiki_url" db:"pad_wiki_url"`
	PadMapUrl                    *string    `json:"pad_map_url" db:"pad_map_url"`
	PadLatitude                  *float64   `json:"pad_latitude" db:"pad_latitude"`
	PadLongitude                 *float64   `json:"pad_longitude" db:"pad_longitude"`
	IdLocation                   *uuid.UUID `json:"id_location" db:"id_location"`
	PadMapImage                  *string    `json:"pad_map_image" db:"pad_map_image"`
	PadTotalLaunchCount          *int32     `json:"pad_total_launch_count" db:"pad_total_launch_count"`
	LocationUrl                  *string    `json:"location_url" db:"location_url"`
	LocationName                 *string    `json:"location_name" db:"location_name"`
	LocationCountryCode          *string    `json:"location_country_code" db:"location_country_code"`
	LocationMapImage             *string    `json:"location_map_image" db:"location_map_image"`
	LocationTotalLaunchCount     *int32     `json:"location_total_launch_count" db:"location_total_launch_count"`
	LocationTotalLandingCount    *int32     `json:"location_total_landing_count" db:"location_total_landing_count"`
}
