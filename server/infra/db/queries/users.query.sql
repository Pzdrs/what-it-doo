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
        username,
        email,
        hashed_password,
        avatar_url,
        bio
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    *;

-- name: GetUserByUsernameOrEmail :one
SELECT
    *
FROM
    users
WHERE
    username = $1
    OR email = $1;
