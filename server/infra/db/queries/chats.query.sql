-- name: ListChats :many
SELECT
    *
FROM
    chats;
    
-- name: GetChatsForUser :many
SELECT
    c.*
FROM
    chats c
JOIN
    chat_participants cp ON c.id = cp.chat_id
WHERE
    cp.user_id = $1;

-- name: GetChatsForUserWithParticipants :many
SELECT
    c.*,
    json_agg(json_build_object('id', u.id, 'name', u.name, 'email', u.email)) AS participants
FROM
    chats c
JOIN
    chat_participants cp ON c.id = cp.chat_id
JOIN
    users u ON cp.user_id = u.id
WHERE
    cp.user_id = $1
GROUP BY
    c.id;

-- name: GetChatParticipants :many
SELECT
    u.*
FROM
    users u
JOIN
    chat_participants cp ON u.id = cp.user_id
WHERE
    cp.chat_id = $1;

-- name: GetChatById :one
SELECT
    *
FROM
    chats
WHERE
    id = $1;

-- name: GetChatByIdWithParticipants :one
SELECT
    c.*,
    json_agg(json_build_object('id', u.id, 'name', u.name, 'email', u.email)) AS participants
FROM
    chats c
JOIN
    chat_participants cp ON c.id = cp.chat_id
JOIN
    users u ON cp.user_id = u.id
WHERE
    c.id = $1
GROUP BY
    c.id;

-- name: GetMessagesForChat :many
SELECT
    *
FROM
   messages 
WHERE
    chat_id = $1
    AND created_at < $3
ORDER BY
    created_at DESC
LIMIT
    $2;

-- name: CreateChat :one
INSERT INTO chats DEFAULT VALUES
RETURNING
    *;

-- name: AddChatParticipant :exec
INSERT INTO chat_participants (chat_id, user_id)
VALUES ($1, $2);

-- name: CreateMessage :one
INSERT INTO messages (chat_id, sender_id, content)
VALUES ($1, $2, $3)
RETURNING
    *;