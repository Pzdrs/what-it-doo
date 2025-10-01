-- name: CreateSession :one
INSERT INTO sessions (user_id, token, device_type, device_os, expires_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;