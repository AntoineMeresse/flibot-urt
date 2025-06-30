-- name: CreatePen :one
INSERT INTO pen (guid, date, size)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPenByDate :many
SELECT * FROM pen
WHERE date = $1
ORDER BY id
LIMIT $2;

-- name: GetPensOrderBySizeAsc :many
SELECT * FROM pen
ORDER BY size ASC
LIMIT $1;

-- name: GetPensOrderBySizeDesc :many
SELECT * FROM pen
ORDER BY size DESC
LIMIT $1;