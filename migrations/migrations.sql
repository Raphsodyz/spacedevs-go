SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;
SET default_tablespace = '';
SET default_table_access_method = heap;

CREATE DATABASE spacedevs;
CREATE ROLE spacedevs_user
WITH LOGIN
PASSWORD 'Spacedevs123456';

GRANT CONNECT ON DATABASE spacedevs TO spacedevs_user;

CREATE SCHEMA data;
CREATE SCHEMA app;

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA data;
COMMENT ON EXTENSION pg_trgm IS 'Text similarity measurement and index searching based on trigrams.';

GRANT USAGE ON SCHEMA data TO spacedevs_user;
GRANT USAGE ON SCHEMA app TO spacedevs_user;

CREATE TABLE IF NOT EXISTS data.orbit(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    name VARCHAR(360) NULL,
    abbrev VARCHAR(360) NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL
);

CREATE TABLE IF NOT EXISTS data.mission(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    launch_library_id INT NULL,
    name VARCHAR(360) NULL,
    description VARCHAR(5000) NULL,
    type VARCHAR(360) NULL,
    id_orbit BIGINT NULL,
    launch_designator VARCHAR(360) NULL,
    search VARCHAR(360) GENERATED ALWAYS AS (
        LOWER(name)
    ) STORED NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL,

    CONSTRAINT fk_mission_orbit FOREIGN KEY (id_orbit) REFERENCES data.orbit(id)
);

CREATE TABLE IF NOT EXISTS data.configuration(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    launch_library_id INT NULL,
    url VARCHAR(1000) NULL,
    name VARCHAR(360) NULL,
    family VARCHAR(360) NULL,
    full_name VARCHAR(360) NULL,
    variant VARCHAR(360) NULL,
    search VARCHAR(600) GENERATED ALWAYS AS (
        LOWER(name || ' ' || family)
    ) STORED NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL
);

CREATE TABLE IF NOT EXISTS data.status(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    name VARCHAR(360) NULL,
    abbrev VARCHAR(360) NULL,
    description VARCHAR(5000) NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL
);

CREATE TABLE IF NOT EXISTS data.launch_service_provider(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    url VARCHAR(1000) NULL,
    name VARCHAR(360) NULL,
    type VARCHAR(360) NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL
);

CREATE TABLE IF NOT EXISTS data.location(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    url VARCHAR(1000) NULL,
    name VARCHAR(360) NULL,
    country_code VARCHAR(360) NULL,
    map_image VARCHAR(1000) NULL,
    total_launch_count INT NULL,
    total_landing_count INT NULL,
    search VARCHAR(360) GENERATED ALWAYS AS (
        LOWER(name)
    ) STORED NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL
);

CREATE TABLE IF NOT EXISTS data.pad(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    url VARCHAR(1000) NULL,
    agency_id INT NULL,
    name VARCHAR(360) NULL,
    info_url VARCHAR(1000) NULL,
    wiki_url VARCHAR(1000) NULL,
    map_url VARCHAR(1000) NULL,
    latitude DOUBLE PRECISION NULL,
    longitude DOUBLE PRECISION NULL,
    id_location BIGINT NULL,
    map_image VARCHAR(1000) NULL,
    total_launch_count INT NULL,
    search VARCHAR(360) GENERATED ALWAYS AS (
        LOWER(name)
    ) STORED NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL,

    CONSTRAINT fk_pad_location FOREIGN KEY (id_location) REFERENCES data.location(id)
);

CREATE TABLE IF NOT EXISTS data.rocket(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    id_configuration BIGINT NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL,

    CONSTRAINT fk_rocket_configuration FOREIGN KEY (id_configuration) REFERENCES data.configuration(id)
);

CREATE TABLE IF NOT EXISTS data.launch(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    api_guid UUID NOT NULL,
    url VARCHAR(1000) NOT NULL,
    launch_library_id INT NULL,
    slug VARCHAR(360) NULL,
    name VARCHAR(360) NULL,
    id_status BIGINT NULL,
    net TIMESTAMP WITHOUT TIME ZONE NULL,
    window_end TIMESTAMP WITHOUT TIME ZONE NULL,
    window_start TIMESTAMP WITHOUT TIME ZONE NULL,
    inhold BOOLEAN NULL,
    tbd_time BOOLEAN NULL,
    tbd_date BOOLEAN NULL,
    probability INT NULL,
    hold_reason VARCHAR(360) NULL,
    fail_reason VARCHAR(750) NULL,
    hashtag VARCHAR(360) NULL,
    id_launch_service_provider BIGINT NULL,
    id_rocket BIGINT NULL,
    id_mission BIGINT NULL,
    id_pad BIGINT NULL,
    web_cast_live BOOLEAN NULL,
    image VARCHAR(1000) NULL,
    infographic VARCHAR(360) NULL,
    programs VARCHAR(360) NULL,
    search VARCHAR(600) GENERATED ALWAYS AS (
        LOWER(name || ' ' || slug)
    ) STORED NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL,

    CONSTRAINT fk_launch_status FOREIGN KEY (id_status) REFERENCES data.status(id),
    CONSTRAINT fk_launch_launch_service_provider FOREIGN KEY (id_launch_service_provider) REFERENCES data.launch_service_provider(id),
    CONSTRAINT fk_launch_rocket FOREIGN KEY (id_rocket) REFERENCES data.rocket(id),
    CONSTRAINT fk_launch_mission FOREIGN KEY (id_mission) REFERENCES data.mission(id),
    CONSTRAINT fk_launch_pad FOREIGN KEY (id_pad) REFERENCES data.pad(id)
);

CREATE MATERIALIZED VIEW IF NOT EXISTS data.launch_view AS
    SELECT
        l.id AS launch_id,
        l.status AS launch_status,
        l.api_guid AS launch_api_guid,
        l.url AS launch_url,
        l.launch_library_id AS launch_launch_library_id,
        l.slug AS launch_slug,
        l.name AS launch_name,
        l.id_status, l.net AS launch_net,
        l.window_end AS launch_window_end,
        l.window_start AS launch_window_start,
        l.inhold AS launch_inhold,
        l.tbd_time AS launch_tbd_time,
        l.tbd_date AS launch_tbd_date,
        l.probability AS launch_probability,
        l.hold_reason AS launch_hold_reason,
        l.fail_reason AS launch_fail_reason,
        l.hashtag AS launch_hashtag,
        l.id_launch_service_provider,
        l.id_rocket,
        l.id_mission,
        l.id_pad,
        l.web_cast_live AS launch_web_cast_live,
        l.image AS launch_image,
        l.infographic AS launch_infographic,
        l.programs AS launch_programs,
        s.name AS status_name,
        s.abbrev AS status_abbrev,
        s.description AS status_description,
        lsp.url AS launch_service_provider_url,
        lsp.name AS launch_service_provider_name,
        lsp.type AS launch_service_provider_type,
        r.id_configuration,
        c.launch_library_id AS configuration_launch_library_id,
        c.url AS configuration_url,
        c.name AS configuration_name,
        c.family AS configuration_family,
        c.full_name AS configuration_full_name,
        c.variant AS configuration_variant,
        m.launch_library_id AS mission_launch_library_id, 
        m.name AS mission_name,
        m.description AS mission_description,
        m.type AS mission_type,
        m.id_orbit,
        m.launch_designator AS mission_launch_designator,
        o.name AS orbit_name,
        o.abbrev AS orbit_abbrev, 
        p.url AS pad_url,
        p.agency_id AS pad_agency_id,
        p.name AS pad_name,
        p.info_url AS pad_info_url,
        p.wiki_url AS pad_wiki_url,
        p.map_url AS pad_map_url,
        p.latitude AS pad_latitude,
        p.longitude AS pad_longitude,
        p.id_location,
        p.map_image AS pad_map_image,
        p.total_launch_count AS pad_total_launch_count,
        loc.url AS location_url,
        loc.name AS location_name,
        loc.country_code AS location_country_code,
        loc.map_image AS location_map_image,
        loc.total_launch_count AS location_total_launch_count,
        loc.total_landing_count AS location_total_landing_count
    FROM
        data.launch AS l
        LEFT JOIN data.status AS s ON l.id_status = s.id AND s.status = 'PUBLISHED' AND s.effective_date = DATE '9999-12-31'
        LEFT JOIN data.launch_service_provider AS lsp ON l.id_launch_service_provider = lsp.id AND l.status = 'PUBLISHED' AND l.effective_date = DATE '9999-12-31'
        LEFT JOIN data.rocket AS r ON l.id_rocket = r.id AND l.status = 'PUBLISHED' AND l.effective_date = DATE '9999-12-31'
        LEFT JOIN data.configuration AS c ON r.id_configuration = c.id AND r.status = 'PUBLISHED' AND r.effective_date = DATE '9999-12-31'
        LEFT JOIN data.mission AS m ON l.id_mission = m.id AND m.status = 'PUBLISHED' AND m.effective_date = DATE '9999-12-31'
        LEFT JOIN data.orbit AS o ON m.id_orbit = o.id AND o.status = 'PUBLISHED' AND o.effective_date = DATE '9999-12-31'
        LEFT JOIN data.pad AS p ON l.id_pad = p.id AND p.status = 'PUBLISHED' AND p.effective_date = DATE '9999-12-31'
        LEFT JOIN data.location AS loc ON p.id_location = loc.id AND loc.status = 'PUBLISHED' AND loc.effective_date = DATE '9999-12-31'
    WHERE
        l.status = 'PUBLISHED'
        AND l.effective_date = DATE '9999-12-31';

CREATE INDEX IDX_GIST_LAUNCH_SLUG_NAME ON data.launch USING gist (search data.gist_trgm_ops);
CREATE INDEX IDX_GIST_CONFIGURATION_NAME_FAMILY ON data.configuration USING gist (search data.gist_trgm_ops);
CREATE INDEX IDX_GIST_MISSION_NAME ON data.mission USING gist (search data.gist_trgm_ops);
CREATE INDEX IDX_GIST_LOCATION_NAME ON data.location USING gist (search data.gist_trgm_ops);
CREATE INDEX IDX_GIST_PAD_NAME ON data.pad USING gist (search data.gist_trgm_ops);

CREATE INDEX idx_mission_id_orbit ON data.mission USING btree(id_orbit);
CREATE INDEX idx_pad_id_location ON data.pad USING btree(id_location);
CREATE INDEX idx_rocket_id_configuration ON data.rocket USING btree(id_configuration);
CREATE INDEX idx_launch_id_status ON data.launch USING btree(id_status);
CREATE INDEX idx_launch_id_launch_service_provider ON data.launch USING btree(id_launch_service_provider);
CREATE INDEX idx_launch_id_rocket ON data.launch USING btree(id_rocket);
CREATE INDEX idx_launch_id_mission ON data.launch USING btree(id_mission);
CREATE INDEX idx_launch_id_pad ON data.launch USING btree(id_pad);

CREATE INDEX idx_configuration_effective_date ON data.configuration USING btree(effective_date);
CREATE INDEX idx_launch_service_provider_effective_date ON data.launch_service_provider USING btree(effective_date);
CREATE INDEX idx_launch_effective_date ON data.launch USING btree(effective_date);
CREATE INDEX idx_location_effective_date ON data.location USING btree(effective_date);
CREATE INDEX idx_mission_effective_date ON data.mission USING btree(effective_date);
CREATE INDEX idx_orbit_effective_date ON data.orbit USING btree(effective_date);
CREATE INDEX idx_pad_effective_date ON data.pad USING btree(effective_date);
CREATE INDEX idx_rocket_effective_date ON data.rocket USING btree(effective_date);
CREATE INDEX idx_status_effective_date ON data.status USING btree(effective_date);

GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES IN SCHEMA data
TO spacedevs_user;

GRANT USAGE, SELECT
ON ALL SEQUENCES IN SCHEMA data
TO spacedevs_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA data
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO spacedevs_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA data
GRANT USAGE, SELECT ON SEQUENCES TO spacedevs_user;

ALTER ROLE spacedevs_user
SET search_path TO data, app, public;