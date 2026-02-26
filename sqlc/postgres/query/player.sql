-- name: CreatePlayer :one
INSERT INTO player (guid, role, name, ip_address, time_joined, aliases)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdatePlayer :one
UPDATE player
SET role = $2,
    name = $3,
    ip_address = $4,
    time_joined = $5,
    aliases = $6
WHERE guid = $1
RETURNING *;

-- name: GetPLayerByGuid :one
SELECT * FROM player
WHERE guid = $1 LIMIT 1;

-- name: ListPlayers :many
SELECT * FROM player
ORDER BY id;

-- name: DeletePlayer :exec
DELETE FROM player
WHERE guid = $1;

-- name: UpdatePlayerOnJoin :exec
UPDATE player
SET name = $2,
    ip_address = $3,
    time_joined = $4,
    aliases = $5
WHERE guid = $1;

-- name: SetPlayerRole :exec
UPDATE player SET role = $2 WHERE guid = $1;

-- name: GetPlayerById :one
SELECT * FROM player
WHERE id = $1 LIMIT 1;

-- name: GetPlayersByIp :many
SELECT * FROM player
WHERE ip_address = $1
ORDER BY name;

-- name: LookupPlayersByNameOrAlias :many
SELECT * FROM player
WHERE name ILIKE '%' || $1 || '%' OR aliases ILIKE '%' || $1 || '%'
ORDER BY name
LIMIT $2;