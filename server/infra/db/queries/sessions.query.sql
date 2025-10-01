-- name: CreateSession :one
INSERT INTO sessions (user_id, token, device_type, device_os, expires_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSessionByToken :one
SELECT * FROM sessions WHERE token = $1;

-- name: GetUserSessions :many
SELECT * FROM sessions WHERE user_id = $1 ORDER BY created_at DESC;