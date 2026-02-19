-- name: UpsertPen :one
INSERT INTO pen (guid, date, size, attempts)
VALUES ($1, $2, $3, 1)
ON CONFLICT (guid, date) DO UPDATE SET
    size = EXCLUDED.size,
    attempts = pen.attempts + 1
RETURNING id, guid, date, size, attempts;

-- name: GetYearlyAttempts :one
SELECT COALESCE(SUM(attempts), 0)::integer
FROM pen
WHERE guid = $1 AND EXTRACT(YEAR FROM date) = EXTRACT(YEAR FROM $2::date);

-- name: GetPlayerPenByDate :one
SELECT size
FROM pen
WHERE guid = $1 AND date = $2;

-- name: DecrementPenAttempts :execrows
UPDATE pen SET attempts = GREATEST(attempts - 1, 0)
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
