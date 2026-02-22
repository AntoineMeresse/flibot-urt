-- name: GetRandomQuote :one
SELECT * FROM quotes
ORDER BY RANDOM()
LIMIT 1;

-- name: SaveQuote :one
INSERT INTO quotes (
    text
) VALUES (
    $1
)
RETURNING *;
