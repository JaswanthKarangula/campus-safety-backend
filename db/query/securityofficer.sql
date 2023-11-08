
-- name: GetAllIssuesByAllSecurityOfficers :many
SELECT i.*
FROM "Issues" i
         JOIN "SecurityOfficerIssues" soi ON i.issue_id = soi.issue_id
         JOIN "SecuritOfficers" so ON soi.officer_id = so.officer_id
         JOIN "Customers" c ON so.customer_id = c.customer_id
WHERE c.customer_id = $1;

-- name: CreateNewSecurityOfficer :one
WITH new_security_officer AS (
INSERT INTO "Users" (username, email, hashedpassword, role)
VALUES ($1, $2, $3, 'Security Officer')
    RETURNING *
    )
INSERT INTO "SecuritOfficers" (officer_id, customer_id)
SELECT user_id, $4
FROM new_security_officer
RETURNING *;

-- name: CreateSecurityOfficerIssue :one
WITH new_issue AS (
INSERT INTO "Issues" (description, status, comments)
VALUES ($1, 'New', '')
    RETURNING issue_id
    )
INSERT INTO "SecurityOfficerIssues" (officer_id, issue_id)
SELECT $2, issue_id
FROM new_issue
RETURNING *;


-- name: UpdateSchedule :one
UPDATE "SecurityOfficerSchedule"
SET start_time = $2,
    end_time = $3,
    day = $4
WHERE officer_id = $1
RETURNING *;

-- name: DeleteOfficerSchedule :many
DELETE FROM "SecurityOfficerSchedule"
WHERE officer_id = $1
RETURNING *;

