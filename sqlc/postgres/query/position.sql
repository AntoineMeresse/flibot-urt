-- name: CreatePosition :one
INSERT INTO position (mapname, location, x, y, z, angle)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPositionByMapAndLocation :one
SELECT * FROM position
WHERE mapname = $1 AND location = $2;

-- name: ListPositionsByMap :many
SELECT * FROM position
WHERE mapname = $1
ORDER BY location ASC;

-- name: DeletePosition :exec
DELETE FROM position
WHERE mapname = $1 AND location = $2;
