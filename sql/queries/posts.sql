-- name: CreatePost :exec
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at, feed_id) 
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);


-- name: GetPostsByUser :many
WITH user_feeds AS(
    SELECT feed_follows.feed_id as feed_id, feeds.name as feed_name
    FROM feed_follows
    JOIN feeds ON feed_follows.feed_id = feeds.id
    WHERE feed_follows.user_id = $1
)
SELECT posts.*, user_feeds.feed_name 
FROM posts
JOIN user_feeds ON posts.feed_id = user_feeds.feed_id
ORDER BY posts.published_at DESC
LIMIT $2;


-- name: DeletePosts :exec
DELETE FROM posts;