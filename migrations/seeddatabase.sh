#!/bin/bash
set -e

migrations_path="/migrations"
seed_path="/migrations/seed"
seed_file="spacedevs_data.sql"

echo "Waiting for PostgreSQL to be ready..."

until pg_isready -U "$POSTGRES_USER" -d "$POSTGRES_DB" -q; do
    sleep 1
done

echo "PostgreSQL is ready."

SCHEMA_EXISTS=$(psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -tAc \
    "SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = 'data';")

if [ "$SCHEMA_EXISTS" -gt "0" ]; then
    echo "Database already initialized, skipping migrations."
    exit 0
fi

if [ -f "$seed_path/$seed_file" ]; then
    echo "Found a backup file, recovering from pg_dump seed..."
    psql -v ON_ERROR_STOP=1 \
         --username "$POSTGRES_USER" \
         --dbname   "$POSTGRES_DB" \
         < "$seed_path/$seed_file"
    echo "Database recovered from seed dump."
else
    echo "Running migration file..."
    psql -v ON_ERROR_STOP=1 \
         --username "$POSTGRES_USER" \
         --dbname   "$POSTGRES_DB" \
         < "$migrations_path/migrations.sql"
    echo "Migrations completed successfully."
fi