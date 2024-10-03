-- name: UserNewAndParse :one
INSERT INTO users (social_provider_user_id, username)
    VALUES (@social_provider_user_id, @username)
RETURNING id;

-- name: InsertStatsInfo :exec
INSERT INTO lc_stats (user_id, easy_submits, medium_submits, hard_submits, total_submits, rank, created_at, updated_at)
       VALUES (@user_id, @easy_submits, @medium_submits, @hard_submits, @total_submits, @rank, @created_at, @updated_at);

-- name: UserGetBySocialProviderId :one
SELECT id FROM users
WHERE social_provider_user_id = @user_slug;

-- name: UserGetStatsByID :one
SELECT u.id,
       u.social_provider_user_id AS userSlug,
       u.username,
       s.easy_submits,
       s.medium_submits,
       s.hard_submits,
       s.total_submits,
       s.rank,
       s.updated_at
FROM users u
         INNER JOIN lc_stats s ON u.id = s.user_id
WHERE u.id = @id;

-- name: UpdateLcStats :exec
UPDATE lc_stats
SET easy_submits = @easy_submits,
    medium_submits = @medium_submits,
    hard_submits = @hard_submits,
    total_submits = @total_submits,
    updated_at = @updated_at
WHERE user_id = @user_id;
