-- name: ListUsers :many
SELECT
    *
FROM
    users;

-- name: CreateUser :one
INSERT INTO
    users (
        first_name,
        last_name,
        email,
        hashed_password,
        avatar_url,
        bio
    )
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

-- name: GetUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = $1;

-- name: GetUserByID :one
SELECT
    *
FROM
    users
WHERE
    id = $1;