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

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;
COMMENT ON EXTENSION pg_trgm IS 'Text similarity measurement and index searching based on trigrams.';

CREATE TABLE IF NOT EXISTS public.orbit(
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

CREATE TABLE IF NOT EXISTS public.mission(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    launch_library_id INT NULL,
    name VARCHAR(360) NULL,
    description VARCHAR(5000) NULL,
    type VARCHAR(360) NULL,
    id_orbit UUID NULL,
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

    CONSTRAINT fk_mission_orbit FOREIGN KEY (id_orbit) REFERENCES public.orbit(id)
);

CREATE TABLE IF NOT EXISTS public.configuration(
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

CREATE TABLE IF NOT EXISTS public.status(
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

CREATE TABLE IF NOT EXISTS public.launch_service_provider(
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

CREATE TABLE IF NOT EXISTS public.location(
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

CREATE TABLE IF NOT EXISTS public.pad(
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

    CONSTRAINT fk_pad_location FOREIGN KEY (id_location) REFERENCES public.location(id)
);

CREATE TABLE IF NOT EXISTS public.rocket(
    id BIGINT PRIMARY KEY,
    id_from_api INT NULL,
    id_configuration BIGINT NULL,
    user_inclusion VARCHAR(20) NULL,
    date_inclusion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    user_change VARCHAR(20) NULL,
    date_change TIMESTAMP WITHOUT TIME ZONE NULL,
    effective_date DATE NOT NULL,
    status VARCHAR(15) NOT NULL

    CONSTRAINT fk_rocket_configuration FOREIGN KEY(id_configuration) REFERENCES public.configuration(id)
);

CREATE TABLE IF NOT EXISTS public.launch(
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
    status VARCHAR(15) NOT NULL

    CONSTRAINT fk_launch_status FOREIGN KEY (id_status) REFERENCES public.status(id),
    CONSTRAINT fk_launch_launch_service_provider FOREIGN KEY (id_launch_service_provider) REFERENCES public.launch_service_provider(id),
    CONSTRAINT fk_launch_rocket FOREIGN KEY (id_rocket) REFERENCES public.rocket(id),
    CONSTRAINT fk_launch_mission FOREIGN KEY (id_mission) REFERENCES public.mission(id),
    CONSTRAINT fk_launch_pad FOREIGN KEY (id_pad) REFERENCES public.pad(id)
);

CREATE INDEX IDX_GIST_LAUNCH_SLUG_NAME ON public.launch USING gist (search public.gist_trgm_ops);
CREATE INDEX IDX_GIST_CONFIGURATION_NAME_FAMILY ON public.configuration USING gist (search public.gist_trgm_ops);
CREATE INDEX IDX_GIST_MISSION_NAME ON public.mission USING gist (search public.gist_trgm_ops);
CREATE INDEX IDX_GIST_LOCATION_NAME ON public.location USING gist (search public.gist_trgm_ops);
CREATE INDEX IDX_GIST_PAD_NAME ON public.pad USING gist (search public.gist_trgm_ops);

CREATE INDEX idx_mission_id_orbit ON public.mission USING btree(id_orbit);
CREATE INDEX idx_pad_id_location ON public.pad USING btree(id_location);
CREATE INDEX idx_rocket_id_configuration ON public.rocket USING btree(id_configuration);
CREATE INDEX idx_launch_id_status ON public.launch USING btree(id_status);
CREATE INDEX idx_launch_id_launch_service_provider ON public.launch USING btree(id_launch_service_provider);
CREATE INDEX idx_launch_id_rocket ON public.launch USING btree(id_rocket);
CREATE INDEX idx_launch_id_mission ON public.launch USING btree(id_mission);
CREATE INDEX idx_launch_id_pad ON public.launch USING btree(id_pad);

CREATE INDEX idx_configuration_effective_date ON public.configuration USING btree(effective_date);
CREATE INDEX idx_launch_service_provider_effective_date ON public.launch_service_provider USING btree(effective_date);
CREATE INDEX idx_launch_effective_date ON public.launch USING btree(effective_date);
CREATE INDEX idx_location_effective_date ON public.location USING btree(effective_date);
CREATE INDEX idx_mission_effective_date ON public.mission USING btree(effective_date);
CREATE INDEX idx_orbit_effective_date ON public.orbit USING btree(effective_date);
CREATE INDEX idx_pad_effective_date ON public.pad USING btree(effective_date);
CREATE INDEX idx_rocket_effective_date ON public.rocket USING btree(effective_date);
CREATE INDEX idx_status_effective_date ON public.status USING btree(effective_date);

SET search_path TO "$user", public;