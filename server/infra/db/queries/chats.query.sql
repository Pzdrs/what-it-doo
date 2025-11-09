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