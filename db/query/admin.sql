

-- name: CreateNewAdmin :one
WITH new_admin AS (
INSERT INTO "Users" (username, email, hashedpassword, role)
VALUES ($1, $2, $3, 'Admin')
    RETURNING user_id
    )
INSERT INTO "Admins" (admin_id)
SELECT user_id
FROM new_admin RETURNING *;
