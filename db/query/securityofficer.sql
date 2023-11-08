
-- name: GetAllIssuesByAllSecurityOfficers :many
SELECT i.*
FROM "Issues" i
         JOIN "SecurityOfficerIssues" soi ON i.issue_id = soi.issue_id
         JOIN "SecuritOfficers" so ON soi.officer_id = so.officer_id
         JOIN "Customers" c ON so.customer_id = c.customer_id
WHERE c.customer_id = $1;