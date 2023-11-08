-- name: UpdateIssueStatus :one
UPDATE "Issues"
SET status = $2,
    comments = $3
WHERE issue_id = $1
RETURNING *;
