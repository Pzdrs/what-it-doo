-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chats;
-- +goose StatementEnd
