-- name: CreateSessions :one
INSERT INTO sessions (
id,
userid,
refresh_token,
expires_at,
created_at
) VALUES ($1,$2,$3,$4,$5) RETURNING *;

-- name: GetSessions :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;

-- name: UpdateSessions :one
UPDATE  sessions set isbloced=$2 where id=$1  RETURNING *;
