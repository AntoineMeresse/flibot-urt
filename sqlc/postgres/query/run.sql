-- name: CreateRun :exec
INSERT INTO runs (guid, utj, mapname, way, runtime, checkpoints, run_date, demopath)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetRuntimeByMapWayUTJ :one
SELECT runtime
FROM runs
WHERE guid = $1 AND mapname = $2 AND way = $3 AND utj = $4
LIMIT 1;

-- name: UpdateRunByGuidAndUTJ :exec
UPDATE runs
SET runtime = $1,
    checkpoints = $2,
    run_date = $3
WHERE guid = $4 AND utj = $5;