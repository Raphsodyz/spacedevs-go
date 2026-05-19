package repository

import (
	"context"
	"fmt"
	"time"

	squirrel "github.com/Masterminds/squirrel"
	models "github.com/Raphsodyz/spacedevs-go/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LaunchRepository interface {
	GetById(ctx context.Context, id int64) (*models.LaunchView, error)
	GetIdsBySlugName(ctx context.Context, slugName string) ([]int64, error)
	GetSearchResults(ctx context.Context, search *models.SearchLaunchFilters) (*models.SearchLaunchResponse, error)
}

type launchRepository struct {
	db *pgxpool.Pool
}

func NewLaunchRepository(db *pgxpool.Pool) LaunchRepository {
	return &launchRepository{db: db}
}

func (r *launchRepository) GetById(ctx context.Context, id int64) (*models.LaunchView, error) {
	var pgsql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := pgsql.Select(
		"launch_id",
		"launch_atualization_date",
		"launch_imported_t",
		"launch_status",
		"launch_api_guid",
		"launch_url",
		"launch_launch_library_id",
		"launch_slug",
		"launch_name",
		"id_status",
		"launch_net",
		"launch_window_end",
		"launch_window_start",
		"launch_inhold",
		"launch_tbd_time",
		"launch_tbd_date",
		"launch_probability",
		"launch_hold_reason",
		"launch_fail_reason",
		"launch_hashtag",
		"id_launch_service_provider",
		"id_rocket",
		"id_mission",
		"id_pad",
		"launch_web_cast_live",
		"launch_image",
		"launch_infographic",
		"launch_programs",
		"status_name",
		"status_abbrev",
		"status_description",
		"launch_service_provider_url",
		"launch_service_provider_name",
		"launch_service_provider_type",
		"id_configuration",
		"configuration_launch_library_id",
		"configuration_url",
		"configuration_name",
		"configuration_family",
		"configuration_full_name",
		"configuration_variant",
		"mission_launch_library_id",
		"mission_name",
		"mission_description",
		"mission_type",
		"id_orbit",
		"mission_launch_designator",
		"orbit_name",
		"orbit_abbrev",
		"pad_url",
		"pad_agency_id",
		"pad_name",
		"pad_info_url",
		"pad_wiki_url",
		"pad_map_url",
		"pad_latitude",
		"pad_longitude",
		"id_location",
		"pad_map_image",
		"pad_total_launch_count",
		"location_url",
		"location_name",
		"location_country_code",
		"location_map_image",
		"location_total_launch_count",
		"location_total_landing_count",
	).
		From("data.launch_view").
		Where(squirrel.Eq{"launch_id": id}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetById query build: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetById query exec: %w", err)
	}

	defer rows.Close()

	launch, err := pgx.CollectOneRow(rows, func(row pgx.CollectableRow) (*models.LaunchView, error) {
		var l models.LaunchView

		err := row.Scan(
			&l.Id,
			&l.AtualizationDate,
			&l.ImportedT,
			&l.EntityStatus,
			&l.ApiGuid,
			&l.Url,
			&l.LaunchLibraryId,
			&l.Slug,
			&l.Name,
			&l.IdStatus,
			&l.Net,
			&l.WindowEnd,
			&l.WindowStart,
			&l.Inhold,
			&l.TbdTime,
			&l.TbdDate,
			&l.Probability,
			&l.HoldReason,
			&l.FailReason,
			&l.Hashtag,
			&l.IdLaunchServiceProvider,
			&l.IdRocket,
			&l.IdMission,
			&l.IdPad,
			&l.WebcastLive,
			&l.Image,
			&l.Infographic,
			&l.Programs,
			&l.StatusName,
			&l.StatusAbbrev,
			&l.StatusDescription,
			&l.LaunchServiceProviderUrl,
			&l.LaunchServiceProviderName,
			&l.LaunchServiceProviderType,
			&l.IdConfiguration,
			&l.ConfigurationLaunchLibraryId,
			&l.ConfigurationUrl,
			&l.ConfigurationName,
			&l.ConfigurationFamily,
			&l.ConfigurationFullName,
			&l.ConfigurationVariant,
			&l.MissionLaunchLibraryId,
			&l.MissionName,
			&l.MissionDescription,
			&l.MissionType,
			&l.IdOrbit,
			&l.MissionLaunchDesignator,
			&l.OrbitName,
			&l.OrbitAbbrev,
			&l.PadUrl,
			&l.PadAgencyId,
			&l.PadName,
			&l.PadInfoUrl,
			&l.PadWikiUrl,
			&l.PadMapUrl,
			&l.PadLatitude,
			&l.PadLongitude,
			&l.IdLocation,
			&l.PadMapImage,
			&l.PadTotalLaunchCount,
			&l.LocationUrl,
			&l.LocationName,
			&l.LocationCountryCode,
			&l.LocationMapImage,
			&l.LocationTotalLaunchCount,
			&l.LocationTotalLandingCount,
		)

		if err != nil {
			return nil, fmt.Errorf("launchRepository.GetById scan: %w", err)
		}

		return &l, nil
	})

	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetById collect: %w", err)
	}

	return launch, nil
}

func (r *launchRepository) GetIdsBySlugName(ctx context.Context, slugName string) ([]int64, error) {
	var pgsql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := pgsql.Select("l.id").
		From("data.launch AS l").
		Where(squirrel.ILike{"l.search": "%" + slugName + "%"}).
		Where(squirrel.Eq{"l.status": "PUBLISHED"}).
		Where(squirrel.Eq{"l.effective_date": time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetIdsBySlugName query build: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetIdsBySlugName query exec: %w", err)
	}

	defer rows.Close()

	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetIdsBySlugName collect: %w", err)
	}

	return ids, nil
}

func (r *launchRepository) GetSearchResults(ctx context.Context, search *models.SearchLaunchFilters) (*models.SearchLaunchResponse, error) {
	const pageSize = 10
	var pgsql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	countQuery, countArgs, err := pgsql.Select("COUNT(*)").
		From("data.launch_view").
		Where(buildSearchFilters(search)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetSearchResults count build: %w", err)
	}

	var totalEntities int
	err = r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&totalEntities)

	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetSearchResults count exec: %w", err)
	}

	if totalEntities == 0 {
		return &models.SearchLaunchResponse{
			NumberOfEntities: 0,
			Entities:         []models.LaunchView{},
		}, nil
	}

	totalPages := 0
	if totalEntities > 0 {
		totalPages = (totalEntities + pageSize - 1) / pageSize
		if totalEntities%pageSize != 0 {
			totalPages--
		}
	}

	if search.Page > totalPages {
		return nil, fmt.Errorf("launchRepository.GetSearchResults: %w", fmt.Errorf("requested page %d exceeds total pages %d", search.Page, totalPages))
	}

	offset := search.Page * pageSize
	dataQuery, dataArgs, err := pgsql.Select(
		"launch_id",
		"launch_atualization_date",
		"launch_imported_t",
		"launch_status",
		"launch_api_guid",
		"launch_url",
		"launch_launch_library_id",
		"launch_slug",
		"launch_name",
		"id_status",
		"launch_net",
		"launch_window_end",
		"launch_window_start",
		"launch_inhold",
		"launch_tbd_time",
		"launch_tbd_date",
		"launch_probability",
		"launch_hold_reason",
		"launch_fail_reason",
		"launch_hashtag",
		"id_launch_service_provider",
		"id_rocket",
		"id_mission",
		"id_pad",
		"launch_web_cast_live",
		"launch_image",
		"launch_infographic",
		"launch_programs",
		"status_name",
		"status_abbrev",
		"status_description",
		"launch_service_provider_url",
		"launch_service_provider_name",
		"launch_service_provider_type",
		"id_configuration",
		"configuration_launch_library_id",
		"configuration_url",
		"configuration_name",
		"configuration_family",
		"configuration_full_name",
		"configuration_variant",
		"mission_launch_library_id",
		"mission_name",
		"mission_description",
		"mission_type",
		"id_orbit",
		"mission_launch_designator",
		"orbit_name",
		"orbit_abbrev",
		"pad_url",
		"pad_agency_id",
		"pad_name",
		"pad_info_url",
		"pad_wiki_url",
		"pad_map_url",
		"pad_latitude",
		"pad_longitude",
		"id_location",
		"pad_map_image",
		"pad_total_launch_count",
		"location_url",
		"location_name",
		"location_country_code",
		"location_map_image",
		"location_total_launch_count",
		"location_total_landing_count",
	).
		From("data.launch_view").
		Where(buildSearchFilters(search)).
		OrderBy("launch_id").
		Limit(pageSize).
		Offset(uint64(offset)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetSearchResults data build: %w", err)
	}

	rows, err := r.db.Query(ctx, dataQuery, dataArgs...)
	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetSearchResults data exec: %w", err)
	}

	defer rows.Close()

	entities, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.LaunchView, error) {
		var l models.LaunchView

		err := row.Scan(
			&l.Id,
			&l.AtualizationDate,
			&l.ImportedT,
			&l.EntityStatus,
			&l.ApiGuid,
			&l.Url,
			&l.LaunchLibraryId,
			&l.Slug,
			&l.Name,
			&l.IdStatus,
			&l.Net,
			&l.WindowEnd,
			&l.WindowStart,
			&l.Inhold,
			&l.TbdTime,
			&l.TbdDate,
			&l.Probability,
			&l.HoldReason,
			&l.FailReason,
			&l.Hashtag,
			&l.IdLaunchServiceProvider,
			&l.IdRocket,
			&l.IdMission,
			&l.IdPad,
			&l.WebcastLive,
			&l.Image,
			&l.Infographic,
			&l.Programs,
			&l.StatusName,
			&l.StatusAbbrev,
			&l.StatusDescription,
			&l.LaunchServiceProviderUrl,
			&l.LaunchServiceProviderName,
			&l.LaunchServiceProviderType,
			&l.IdConfiguration,
			&l.ConfigurationLaunchLibraryId,
			&l.ConfigurationUrl,
			&l.ConfigurationName,
			&l.ConfigurationFamily,
			&l.ConfigurationFullName,
			&l.ConfigurationVariant,
			&l.MissionLaunchLibraryId,
			&l.MissionName,
			&l.MissionDescription,
			&l.MissionType,
			&l.IdOrbit,
			&l.MissionLaunchDesignator,
			&l.OrbitName,
			&l.OrbitAbbrev,
			&l.PadUrl,
			&l.PadAgencyId,
			&l.PadName,
			&l.PadInfoUrl,
			&l.PadWikiUrl,
			&l.PadMapUrl,
			&l.PadLatitude,
			&l.PadLongitude,
			&l.IdLocation,
			&l.PadMapImage,
			&l.PadTotalLaunchCount,
			&l.LocationUrl,
			&l.LocationName,
			&l.LocationCountryCode,
			&l.LocationMapImage,
			&l.LocationTotalLaunchCount,
			&l.LocationTotalLandingCount,
		)

		if err != nil {
			return l, fmt.Errorf("launchRepository.GetSearchResults scan: %w", err)
		}

		return l, nil
	})

	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetSearchResults collect: %w", err)
	}

	return &models.SearchLaunchResponse{
		Entities:         entities,
		NumberOfPages:    totalPages,
		CurrentPage:      search.Page,
		NumberOfEntities: totalEntities,
	}, nil
}

func buildSearchFilters(search *models.SearchLaunchFilters) squirrel.And {
	conds := squirrel.And{}

	if search == nil {
		return conds
	}

	if len(search.MissionIds) > 0 {
		conds = append(conds, squirrel.Eq{"mission_id": search.MissionIds})
	}
	if len(search.RocketIds) > 0 {
		conds = append(conds, squirrel.Eq{"rocket_id": search.RocketIds})
	}
	if len(search.LocationIds) > 0 {
		conds = append(conds, squirrel.Eq{"location_id": search.LocationIds})
	}
	if len(search.PadIds) > 0 {
		conds = append(conds, squirrel.Eq{"pad_id": search.PadIds})
	}
	if len(search.LaunchIds) > 0 {
		conds = append(conds, squirrel.Eq{"launch_id": search.LaunchIds})
	}

	return conds
}
