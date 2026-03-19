-- name: UpsertServer :exec
INSERT INTO server (ip, port, rconpassword, channel_id, name)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (ip, port) DO UPDATE SET rconpassword = EXCLUDED.rconpassword, name = EXCLUDED.name;
