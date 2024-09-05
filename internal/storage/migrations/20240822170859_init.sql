-- +goose Up
-- +goose StatementBegin
CREATE TYPE SOCIAL_PROVIDER AS ENUM (
    'leetcode'
    );

CREATE TABLE users
(
    id                      BIGSERIAL                NOT NULL PRIMARY KEY,
    social_provider_user_id VARCHAR                  NOT NULL,
    username                VARCHAR                  NOT NULL,
    created_at              TIMESTAMP WITH TIME ZONE NOT NULL,
    UNIQUE (social_provider_user_id)
)

CREATE TABLE lc_stats
(
    user_id                 BIGINT                  NOT NULL PRIMARY KEY REFERENCES users(id)
    easy_submits            INTEGER
    medium_submits          INTEGER
    hard_submits            INTEGER
    rank                    BIGINT                   NOT NULL
    created_at              TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at              TIMESTAMP WITH TIME ZONE NOT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE lc_stats;
DROP TABLE users;

DROP TYPE SOCIAL_PROVIDER;
-- +goose StatementEnd

