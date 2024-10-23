-- +goose Up
-- +goose StatementBegin
CREATE TYPE SOCIAL_PROVIDER AS ENUM (
    'leetcode'
);

CREATE TABLE users (
                       id                      BIGSERIAL PRIMARY KEY NOT NULL,
                       lc_user_id  VARCHAR NOT NULL,
                       UNIQUE (lc_user_id)
);

CREATE TABLE lc_stats (
                          user_id                 BIGINT PRIMARY KEY REFERENCES users(id),
                          username                VARCHAR NOT NULL,
                          easy_submits            BIGINT,
                          medium_submits          BIGINT,
                          hard_submits            BIGINT,
                          total_submits           BIGINT,
                          rank                    BIGINT NOT NULL,
                          created_at              TIMESTAMP WITH TIME ZONE NOT NULL,
                          updated_at              TIMESTAMP WITH TIME ZONE NOT NULL,
                          UNIQUE (user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE lc_stats;
DROP TABLE users;

DROP TYPE SOCIAL_PROVIDER;
-- +goose StatementEnd
