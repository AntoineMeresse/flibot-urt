-- name: UpsertPen :one
INSERT INTO pen (guid, date, size)
VALUES ($1, $2, $3)
ON CONFLICT (guid, date) DO UPDATE SET
    size = EXCLUDED.size
RETURNING id, guid, date, size;

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

-- name: GetPenCounter :one
SELECT attempts
FROM pen_counter
WHERE guid = $1 AND year = $2;

-- name: IncrementPenCounter :exec
INSERT INTO pen_counter (guid, year, attempts)
VALUES ($1, $2, 1)
ON CONFLICT (guid, year) DO UPDATE SET
    attempts = pen_counter.attempts + 1;

-- name: DecrementPenCounter :exec
INSERT INTO pen_counter (guid, year, attempts)
VALUES ($1, $2, -1)
ON CONFLICT (guid, year) DO UPDATE SET
    attempts = pen_counter.attempts - 1;
