-- name: InsertStatsInfo :exec
INSERT INTO lc_stats (user_slug, username, easy_submits, medium_submits, hard_submits, total_submits, rank, created_at, updated_at)
       VALUES (@user_slug, @username, @easy_submits, @medium_submits, @hard_submits, @total_submits, @rank, @created_at, @updated_at);

-- name: UserGetStatsBySlug :one
SELECT user_slug,
       username,
       easy_submits,
       medium_submits,
       hard_submits,
       total_submits,
       rank,
       updated_at
FROM lc_stats
WHERE user_slug = @user_slug;

-- name: UpdateLcStats :exec
UPDATE lc_stats
SET easy_submits = @easy_submits,
    medium_submits = @medium_submits,
    hard_submits = @hard_submits,
    total_submits = @total_submits,
    updated_at = @updated_at
WHERE user_slug = @user_slug;
