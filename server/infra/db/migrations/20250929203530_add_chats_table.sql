-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    chats (
        id BIGSERIAL PRIMARY KEY,
        title VARCHAR(255),
        created_at TIMESTAMPTZ DEFAULT NOW (),
        updated_at TIMESTAMPTZ DEFAULT NOW ()
    );

CREATE TABLE
    chat_participants (
        chat_id BIGINT REFERENCES chats (id) ON DELETE CASCADE,
        user_id UUID REFERENCES users (id) ON DELETE CASCADE,
        joined_at TIMESTAMPTZ DEFAULT NOW (),
        PRIMARY KEY (chat_id, user_id)
    );

CREATE TABLE
    messages (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        chat_id BIGINT REFERENCES chats (id) ON DELETE CASCADE,
        sender_id UUID REFERENCES users (id) ON DELETE SET NULL,
        replying_to UUID REFERENCES messages (id) ON DELETE SET NULL,
        content TEXT,
        created_at TIMESTAMPTZ DEFAULT NOW (),
        delivered_at TIMESTAMPTZ,
        read_at TIMESTAMPTZ,
        updated_at TIMESTAMPTZ DEFAULT NOW ()
    );

CREATE TRIGGER update_chats_updated_at BEFORE
UPDATE ON chats FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE chats;

DROP TABLE chat_participants;

DROP TABLE messages;

DELETE TRIGGER IF EXISTS update_chats_updated_at ON chats;

-- +goose StatementEnd