-- name: AddBan :exec
INSERT INTO bans (guid, ip, reason) VALUES ($1, $2, $3);

-- name: GetBan :one
SELECT reason FROM bans WHERE guid = $1 LIMIT 1;
