-- name: UserNewAndParse :one
INSERT INTO users (lc_user_id)
    VALUES (@lc_user_id)
RETURNING id;

-- name: InsertStatsInfo :exec
INSERT INTO lc_stats (user_id, username, easy_submits, medium_submits, hard_submits, total_submits, rank, created_at, updated_at)
       VALUES (@user_id, @username, @easy_submits, @medium_submits, @hard_submits, @total_submits, @rank, @created_at, @updated_at);

-- name: UserGetByLeetCodeId :one
SELECT id FROM users
WHERE lc_user_id = @user_slug;

-- name: UserGetStatsByID :one
SELECT u.id,
       u.lc_user_id AS userSlug,
       s.username,
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
WHERE user_id = @user_id IS NOT NULL;
