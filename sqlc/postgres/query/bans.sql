-- name: AddBan :exec
INSERT INTO bans (guid, ip, reason) VALUES ($1, $2, $3);

-- name: GetBan :one
SELECT reason FROM bans WHERE guid = $1 LIMIT 1;

-- name: RemoveBan :exec
DELETE FROM bans WHERE guid = (SELECT guid FROM player WHERE id = $1);

-- name: GetBans :many
SELECT p.id, p.name FROM bans b JOIN player p ON p.guid = b.guid;
