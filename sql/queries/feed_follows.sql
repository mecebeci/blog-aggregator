-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, NOW(), NOW(), $2, $3)
    RETURNING *
)
SELECT 
    inserted.id,
    inserted.created_at,
    inserted.updated_at,
    inserted.user_id,
    inserted.feed_id,
    users.name AS user_name,
    feeds.name AS feed_name
FROM inserted
JOIN users ON users.id = inserted.user_id
JOIN feeds ON feeds.id = inserted.feed_id;


-- name: GetFeedFollowsForUser :many
SELECT 
    ff.id AS feed_follow_id,
    ff.created_at,
    ff.updated_at,
    u.id AS user_id,
    u.name AS user_name,
    f.id AS feed_id,
    f.name AS feed_name,
    f.url AS feed_url
FROM feed_follows ff
JOIN users u ON u.id = ff.user_id
JOIN feeds f ON f.id = ff.feed_id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows 
WHERE user_id = $1 AND feed_id = $2; 
