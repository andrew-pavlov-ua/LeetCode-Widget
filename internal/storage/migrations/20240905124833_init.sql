-- +goose Up
-- +goose StatementBegin
CREATE TABLE lc_stats
(
    id             BIGSERIAL                NOT NULL PRIMARY KEY,
    lc_user_slug   VARCHAR                  NOT NULL UNIQUE,
    username       VARCHAR                  NOT NULL,
    easy_submits   BIGINT                   NOT NULL,
    medium_submits BIGINT                   NOT NULL,
    hard_submits   BIGINT                   NOT NULL,
    total_submits  BIGINT                   NOT NULL,
    rank           BIGINT                   NOT NULL,
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at     TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE profile_hourly_visits_stats
(
    user_id BIGINT    NOT NULL REFERENCES lc_stats (id),
    time    TIMESTAMP NOT NULL,
    count   BIGINT    NOT NULL,
    PRIMARY KEY (user_id, time)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE profile_hourly_visits_stats;
DROP TABLE lc_stats;
-- +goose StatementEnd
