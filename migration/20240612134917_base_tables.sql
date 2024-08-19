-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts
(
    id             UUID PRIMARY KEY,
    email          TEXT UNIQUE NOT NULL,
    nickname       TEXT UNIQUE NOT NULL,
    password_hash  TEXT        NOT NULL,
    role_code      TEXT        NOT NULL,
    created_at     TIMESTAMP   NOT NULL,
    last_active_at TIMESTAMP
);

CREATE TABLE devices
(
    id          UUID PRIMARY KEY,
    account_id  UUID REFERENCES accounts (id) NOT NULL,
    name        TEXT                          NOT NULL,
    fingerprint TEXT                          NOT NULL
);

CREATE TABLE refresh_tokens
(
    id         UUID PRIMARY KEY,
    device_id  UUID REFERENCES devices (id) NOT NULL,
    issued_at  TIMESTAMP                    NOT NULL,
    expires_at TIMESTAMP                    NOT NULL,
);

CREATE TABLE login_attempts
(
    id           UUID PRIMARY KEY,
    email        TEXT      NOT NULL,
    success      BOOLEAN   NOT NULL,
    attempted_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE login_attempts;
DROP TABLE refresh_tokens;
DROP TABLE devices;
DROP TABLE accounts;
-- +goose StatementEnd
