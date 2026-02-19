-- name: CreatePen :one
INSERT INTO pen (guid, date, size)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPlayerPenByDate :one
SELECT size
FROM pen
WHERE guid = $1 AND date = $2;

-- name: GetAllPenByDate :many
SELECT pen.*, player.name
FROM pen
JOIN player ON pen.guid = player.guid
WHERE date = $1
ORDER BY size DESC
LIMIT $2;

-- name: GetPensOrderBySizeAsc :many
SELECT pen.*, player.name
FROM pen
JOIN player ON pen.guid = player.guid
WHERE EXTRACT(YEAR FROM date) = EXTRACT(YEAR FROM $1::date)
ORDER BY size ASC
LIMIT $2;

-- name: GetPensOrderBySizeDesc :many
SELECT pen.*, player.name
FROM pen
JOIN player ON pen.guid = player.guid
WHERE EXTRACT(YEAR FROM date) = EXTRACT(YEAR FROM $1::date)
ORDER BY size DESC
LIMIT $2;
