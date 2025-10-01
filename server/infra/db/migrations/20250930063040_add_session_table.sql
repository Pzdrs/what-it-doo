-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    sessions (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        token VARCHAR(255) UNIQUE NOT NULL,
        device_type VARCHAR(255),
        device_os VARCHAR(255),
        revoked_at TIMESTAMPTZ,
        expires_at TIMESTAMPTZ NOT NULL,
        created_at TIMESTAMPTZ DEFAULT NOW (),
        updated_at TIMESTAMPTZ DEFAULT NOW ()
    );

CREATE TRIGGER update_sessions_updated_at BEFORE
UPDATE ON sessions FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;

DELETE TRIGGER IF EXISTS update_sessions_updated_at ON sessions;
-- +goose StatementEnd