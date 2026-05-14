package repository

import (
	"context"
	"fmt"
	"time"

	squirrel "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	entity "github.com/spacedevs-go/internal/entity"
	models "github.com/spacedevs-go/models"
)

type LaunchRepository interface {
	GetById(ctx context.Context, id int64) (*entity.Launch, error)
	GetIdsBySlugName(ctx context.Context, slugName string) ([]int64, error)
	GetSearchResults(ctx context.Context, search *models.SearchLaunchRequest) ([]*entity.Launch, error)
}

type launchRepository struct {
	db *pgxpool.Pool
}

func (r *launchRepository) GetById(ctx context.Context, id int64) (*entity.Launch, error) {
	var pgsql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := pgsql.Select(
		// launch
		"l.id",
		"l.id_from_api",
		"l.api_guid",
		"l.url",
		"l.launch_library_id",
		"l.slug",
		"l.name",
		"l.status",
		"l.id_status",
		"l.net",
		"l.window_end",
		"l.window_start",
		"l.inhold",
		"l.tbd_time",
		"l.tbd_date",
		"l.probability",
		"l.hold_reason",
		"l.fail_reason",
		"l.hashtag",
		"l.id_launch_service_provider",
		"l.id_rocket",
		"l.id_mission",
		"l.id_pad",
		"l.web_cast_live",
		"l.image",
		"l.infographic",
		"l.programs",

		// status
		"s.id",
		"s.id_from_api",
		"s.name",
		"s.abbrev",
		"s.description",

		// launch service provider
		"lsp.id",
		"lsp.id_from_api",
		"lsp.url",
		"lsp.name",
		"lsp.type",

		// rocket
		"r.id",
		"r.id_from_api",
		"r.id_configuration",

		// configuration
		"c.id",
		"c.id_from_api",
		"c.launch_library_id",
		"c.url",
		"c.name",
		"c.family",
		"c.full_name",
		"c.variant",

		// mission
		"m.id",
		"m.id_from_api",
		"m.launch_library_id",
		"m.name",
		"m.description",
		"m.type",
		"m.id_orbit",
		"m.launch_designator",

		// orbit
		"o.id",
		"o.id_from_api",
		"o.name",
		"o.abbrev",

		// pad
		"p.id",
		"p.id_from_api",
		"p.url",
		"p.agency_id",
		"p.name",
		"p.info_url",
		"p.wiki_url",
		"p.map_url",
		"p.latitude",
		"p.longitude",
		"p.id_location",
		"p.map_image",
		"p.total_launch_count",

		// location
		"loc.id",
		"loc.id_from_api",
		"loc.url",
		"loc.name",
		"loc.country_code",
		"loc.map_image",
		"loc.total_launch_count",
		"loc.total_landing_count",
	).
		From("public.launch AS l").
		LeftJoin("public.status                  AS s   ON s.id   = l.id_status").
		LeftJoin("public.launch_service_provider AS lsp ON lsp.id = l.id_launch_service_provider").
		LeftJoin("public.rocket                  AS r   ON r.id   = l.id_rocket").
		LeftJoin("public.configuration           AS c   ON c.id   = r.id_configuration").
		LeftJoin("public.mission                 AS m   ON m.id   = l.id_mission").
		LeftJoin("public.orbit                   AS o   ON o.id   = m.id_orbit").
		LeftJoin("public.pad                     AS p   ON p.id   = l.id_pad").
		LeftJoin("public.location                AS loc ON loc.id = p.id_location").
		Where(squirrel.Eq{"l.id": id}).
		Where(squirrel.Eq{"l.status": "PUBLISHED"}).
		Where(squirrel.Eq{"l.effective_date": time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetById data build: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetById data exec: %w", err)
	}

	defer rows.Close()

	launch, err := pgx.CollectOneRow(rows, func(row pgx.CollectableRow) (*entity.Launch, error) {
		var l entity.Launch

		err := row.Scan(
			// launch
			&l.Id,
			&l.IdFromApi,
			&l.ApiUuid,
			&l.Url,
			&l.LaunchLibraryId,
			&l.Slug,
			&l.Name,
			&l.ProcessingStatus,
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

			// status
			&l.Status.Id,
			&l.Status.IdFromApi,
			&l.Status.Name,
			&l.Status.Abbrev,
			&l.Status.Description,

			// launch service provider
			&l.LaunchServiceProvider.Id,
			&l.LaunchServiceProvider.IdFromApi,
			&l.LaunchServiceProvider.Url,
			&l.LaunchServiceProvider.Name,
			&l.LaunchServiceProvider.Type,

			// rocket
			&l.Rocket.Id,
			&l.Rocket.IdFromApi,
			&l.Rocket.IdConfiguration,

			// configuration
			&l.Rocket.Configuration.Id,
			&l.Rocket.Configuration.IdFromApi,
			&l.Rocket.Configuration.LaunchLibraryId,
			&l.Rocket.Configuration.Url,
			&l.Rocket.Configuration.Name,
			&l.Rocket.Configuration.Family,
			&l.Rocket.Configuration.FullName,
			&l.Rocket.Configuration.Variant,

			// mission
			&l.Mission.Id,
			&l.Mission.IdFromApi,
			&l.Mission.LaunchLibraryId,
			&l.Mission.Name,
			&l.Mission.Description,
			&l.Mission.Type,
			&l.Mission.IdOrbit,
			&l.Mission.LaunchDesignator,

			// orbit
			&l.Mission.Orbit.Id,
			&l.Mission.Orbit.IdFromApi,
			&l.Mission.Orbit.Name,
			&l.Mission.Orbit.Abbrev,

			// pad
			&l.Pad.Id,
			&l.Pad.IdFromApi,
			&l.Pad.Url,
			&l.Pad.AgencyId,
			&l.Pad.Name,
			&l.Pad.InfoUrl,
			&l.Pad.WikiUrl,
			&l.Pad.MapUrl,
			&l.Pad.Latitude,
			&l.Pad.Longitude,
			&l.Pad.IdLocation,
			&l.Pad.MapImage,
			&l.Pad.TotalLaunchCount,

			// location
			&l.Pad.Location.Id,
			&l.Pad.Location.IdFromApi,
			&l.Pad.Location.Url,
			&l.Pad.Location.Name,
			&l.Pad.Location.CountryCode,
			&l.Pad.Location.MapImage,
			&l.Pad.Location.TotalLaunchCount,
			&l.Pad.Location.TotalLandingCount,
		)

		if err != nil {
			return nil, fmt.Errorf("launchRepository.GetById scan: %w", err)
		}

		return &l, nil
	})

	return launch, nil
}

func (r *launchRepository) GetIdsBySlugName(ctx context.Context, slugName string) ([]int64, error) {
	var pgsql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := pgsql.Select(
		"l.id",
	).
		From("public.launch AS l").
		Where(squirrel.ILike{"l.search": slugName}).
		Where(squirrel.Eq{"l.status": "PUBLISHED"}).
		Where(squirrel.Eq{"l.effective_date": time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetIdsBySlugName data build: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetIdsBySlugName data exec: %w", err)
	}

	defer rows.Close()

	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, fmt.Errorf("launchRepository.GetIdsBySlugName collect: %w", err)
	}

	return ids, nil
}

func (r *launchRepository) GetSearchResults(ctx context.Context, search *models.SearchLaunchRequest) ([]*entity.Launch, error) {

}
