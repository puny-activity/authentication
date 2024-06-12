-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts
(
    id              UUID PRIMARY KEY,
    username        TEXT UNIQUE NOT NULL,
    hashed_password TEXT        NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL,
    last_online     TIMESTAMPTZ
);

CREATE TABLE account_roles
(
    account_id UUID REFERENCES accounts (id) NOT NULL,
    role       TEXT                          NOT NULL
);

CREATE TABLE devices
(
    id          UUID PRIMARY KEY,
    account_id  UUID REFERENCES accounts (id) NOT NULL,
    fingerprint TEXT                          NOT NULL
);

CREATE TABLE refresh_tokens
(
    id         UUID PRIMARY KEY,
    device_id  UUID REFERENCES devices (id) NOT NULL,
    issued_at  TIMESTAMPTZ                  NOT NULL,
    expires_at TIMESTAMPTZ                  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_tokens;
DROP TABLE devices;
DROP TABLE account_roles;
DROP TABLE accounts;
-- +goose StatementEnd
