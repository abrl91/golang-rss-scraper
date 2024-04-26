-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, title, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched ASC NULLS FIRST
LIMIT $1;

-- name: UpdateFeedLastFetched :one
UPDATE feeds SET last_fetched = NOW(), updated_at = NOW()
WHERE id = $1
RETURNING *;
