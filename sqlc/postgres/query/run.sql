-- name: CreateRun :exec
INSERT INTO runs (guid, utj, mapname, way, runtime, checkpoints, run_date, demopath)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetRuntimeByMapWayUTJ :one
SELECT runtime
FROM runs
WHERE mapname = $1 AND way = $2 AND utj = $3
ORDER BY runtime ASC
LIMIT 1;

-- name: UpdateRunByGuidAndUTJ :exec
UPDATE runs
SET runtime = $1,
    checkpoints = $2,
    run_date = $3
WHERE guid = $4 AND utj = $5;