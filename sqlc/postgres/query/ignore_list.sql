-- name: AddIgnore :exec
INSERT INTO ignore_list (guid, ignored_guid) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: GetIgnoredGuids :many
SELECT ignored_guid FROM ignore_list WHERE guid = $1;

-- name: GetIgnoredPlayers :many
SELECT p.id, il.ignored_guid, p.name
FROM ignore_list il
JOIN player p ON p.guid = il.ignored_guid
WHERE il.guid = $1
ORDER BY p.id;

-- name: RemoveIgnore :exec
DELETE FROM ignore_list WHERE guid = $1 AND ignored_guid = $2;
