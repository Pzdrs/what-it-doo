-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    chats (
        id BIGSERIAL PRIMARY KEY,
        title VARCHAR(255),
        created_at TIMESTAMPTZ DEFAULT NOW (),
        updated_at TIMESTAMPTZ DEFAULT NOW ()
    );

CREATE TRIGGER update_chats_updated_at BEFORE
UPDATE ON chats FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chats;

DELETE TRIGGER IF EXISTS update_chats_updated_at ON chats;
-- +goose StatementEnd