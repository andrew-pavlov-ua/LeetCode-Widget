-- +goose Up
-- +goose StatementBegin
CREATE TABLE lc_stats (
                          lc_user_id              VARCHAR NOT NULL PRIMARY KEY,
                          username                VARCHAR NOT NULL,
                          easy_submits            BIGINT NOT NULL,
                          medium_submits          BIGINT NOT NULL,
                          hard_submits            BIGINT NOT NULL,
                          total_submits           BIGINT NOT NULL,
                          rank                    BIGINT NOT NULL,
                          created_at              TIMESTAMP WITH TIME ZONE NOT NULL,
                          updated_at              TIMESTAMP WITH TIME ZONE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE lc_stats;

-- +goose StatementEnd
