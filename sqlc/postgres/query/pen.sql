-- name: CreatePen :one
INSERT INTO pen (guid, date, size)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPenByDate :many
SELECT pen.*, player.name
FROM pen
JOIN player ON pen.guid = player.guid
WHERE date = $1
ORDER BY size ASC
LIMIT $2;

-- name: GetPensOrderBySizeAsc :many
SELECT pen.*, player.name
FROM pen
JOIN player ON pen.guid = player.guid
ORDER BY size ASC
LIMIT $1;

-- name: GetPensOrderBySizeDesc :many
SELECT pen.*, player.name
FROM pen
JOIN player ON pen.guid = player.guid
ORDER BY size DESC
LIMIT $1;