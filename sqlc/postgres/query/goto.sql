-- name: GetGotoByMapAndJump :one
SELECT pos_x, pos_y, pos_z, angle_v, angle_h
FROM goto
WHERE mapname = $1 AND jumpname = $2;

-- name: GetGotoNamesByMap :many
SELECT jumpname FROM goto WHERE mapname = $1 ORDER BY jumpname;

-- name: DeleteGoto :execrows
DELETE FROM goto WHERE mapname = $1 AND jumpname = $2;

-- name: UpsertGoto :exec
INSERT INTO goto (mapname, jumpname, pos_x, pos_y, pos_z, angle_v, angle_h)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (mapname, jumpname) DO UPDATE SET
    pos_x = EXCLUDED.pos_x,
    pos_y = EXCLUDED.pos_y,
    pos_z = EXCLUDED.pos_z,
    angle_v = EXCLUDED.angle_v,
    angle_h = EXCLUDED.angle_h;
