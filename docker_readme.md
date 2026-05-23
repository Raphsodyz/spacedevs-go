# Tail logs
docker compose logs -f api

# Rebuild the API image after code changes
docker compose up -d --build api

# Connect to postgres
docker exec -it spacedevs-postgresql psql -U spacedevs_user -d spacedevs

# Wipe volumes and start fresh
docker compose down -v && docker compose up -d