-- name: UserNewAndParse :one
INSERT INTO users (social_provider_user_id, username)
    VALUES (@social_provider_user_id, @username)
RETURNING id;

-- name: InsertStatsInfo :exec
INSERT INTO lc_stats (user_id, easy_submits, medium_submits, hard_submits, total_submits, created_at, updated_at)
       VALUES (@user_id, @easy_submits, @medium_submits, @hard_submits, @total_submits, @created_at, @updated_at);