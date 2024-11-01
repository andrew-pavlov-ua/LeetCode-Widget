-- name: ProfileHourlyVisitsStatsUpsert :exec
INSERT INTO profile_hourly_visits_stats (user_slug, time, count)
VALUES (@user_slug, @time, @count)
ON CONFLICT (user_slug, time) DO UPDATE
    SET count = profile_hourly_visits_stats.count + @count;

-- name: TotalCount :one
SELECT COALESCE(SUM(count), 0)::BIGINT as count 
FROM profile_hourly_visits_stats 
WHERE user_slug = @user_slug;

-- name: ProfileVisitsStatsByPeriod :one
SELECT COALESCE(SUM(count), 0)::BIGINT as count
FROM profile_hourly_visits_stats
WHERE user_slug = @user_slug
    AND time >= @start_time;

-- name: ProfileHourlyViewsStats :one
SELECT COALESCE(SUM(count) FILTER ( WHERE time >= @day ), 0)::BIGINT  AS day_count,
       COALESCE(SUM(count) FILTER ( WHERE time >= @week ), 0)::BIGINT AS week_count,
       SUM(count)                                                     AS month_count
FROM profile_hourly_visits_stats
WHERE user_slug = @user_slug
  AND time >= @month
GROUP BY user_slug;