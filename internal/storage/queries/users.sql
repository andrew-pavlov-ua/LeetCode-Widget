-- name UserNewAndParse :id
INSERT INTO users (social_provider_user_id, username, created_at, updated_at)
    VALUES (@social_provider_user_id, @username, @created_at, @updated_at)
INSERT INTO lc_stats (user_id, easy_submits, medium_submits, hard_submits total_submits, created_at, updated_at)
    VALUES (id, @easy_submits, @medium_submits, @hard_submits @total_submits, @created_at, @updated_at)
RETURNING id;

