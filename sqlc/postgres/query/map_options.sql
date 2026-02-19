-- name: UpsertMapOptions :exec
INSERT INTO map_options (mapname, options)
VALUES ($1, $2)
ON CONFLICT (mapname) DO UPDATE SET options = EXCLUDED.options;

-- name: GetMapOptions :one
SELECT options FROM map_options WHERE mapname = $1;

-- name: DeleteMapOptions :execrows
DELETE FROM map_options WHERE mapname = $1;
