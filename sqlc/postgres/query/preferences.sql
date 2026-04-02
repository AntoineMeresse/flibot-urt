-- name: UpsertPreferences :exec
INSERT INTO player_preferences (guid, commands)
VALUES ($1, $2)
ON CONFLICT (guid) DO UPDATE SET commands = $2;

-- name: GetPreferences :one
SELECT commands FROM player_preferences WHERE guid = $1;
