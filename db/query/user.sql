
-- name: GetUser :one
SELECT *
FROM "Users"
WHERE user_id = $1;
