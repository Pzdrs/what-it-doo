-- +goose Up
-- +goose StatementBegin
ALTER TABLE users 
ADD is_online BOOLEAN DEFAULT FALSE,
ADD last_active_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users 
DROP COLUMN IF EXISTS is_online,
DROP COLUMN IF EXISTS last_active_at;
-- +goose StatementEnd
