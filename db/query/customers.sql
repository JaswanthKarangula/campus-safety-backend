-- name: CreateNewCustomer :one
WITH new_customer AS (
INSERT INTO "Users" (username, email, hashedpassword, role)
VALUES ($1, $2, $3, 'Customer')
    RETURNING *
    )
INSERT INTO "Customers" (customer_id, customer_name, admin_id)
SELECT user_id, username, 1
FROM new_customer
    RETURNING *;


-- name: DeleteDrone :one
DELETE FROM "Drones"
WHERE drone_id = $1 RETURNING *;




-- name: GetAllIssuesByCustomer :many
SELECT i.*
FROM "Issues" i
         JOIN "Customers" c ON i.customer_id = c.customer_id
WHERE c.customer_id = $1;

-- name: GetAllDronesByCustomer :many
SELECT d.*
FROM "Drones" d
         JOIN "Customers" c ON d.customer_id = c.customer_id
WHERE c.customer_id = $1;

-- name: GetAllActiveSecurityOfficers :many
SELECT u.*
FROM "Users" u
         JOIN "SecuritOfficers" so ON u.user_id = so.officer_id
         JOIN "SecurityOfficerSchedule" sch ON so.officer_id = sch.officer_id
         JOIN "Customers" c ON so.customer_id = c.customer_id
WHERE c.customer_id = $1
  AND sch.day = TO_CHAR(NOW(), 'Dy')
  AND sch.start_time <= TO_CHAR(NOW(), 'HH24:MI:SS')::time
  AND sch.end_time >= TO_CHAR(NOW(), 'HH24:MI:SS')::time;




-- name: AddNewDrone :one
INSERT INTO "Drones" (model, customer_id, status, last_active)
VALUES ($1,$2,$3, NOW())
    RETURNING *;


-- name: StopDroneStream :one
UPDATE "Feeds"
SET status = 'inactive'
WHERE drone_id = $1 and stream_id=$2 RETURNING *;

-- name: GetAllAlerts :one
SELECT sda.*
FROM "SafetyDetectionAlerts" sda
         JOIN "Drones" d ON sda.drone_id = d.drone_id
         JOIN "Customers" c ON d.customer_id = c.customer_id
WHERE c.customer_id = $1;

-- name: GetAlertsCountsByPriority :many
SELECT sda.detection_type_id, COUNT(*) as alert_count
FROM "SafetyDetectionAlerts" sda
         JOIN "Drones" d ON sda.drone_id = d.drone_id
         JOIN "Customers" c ON d.customer_id = c.customer_id
WHERE c.customer_id = $1
GROUP BY sda.detection_type_id
ORDER BY sda.detection_type_id;

-- name: GetUnResolvedAlertsCountsByPriority :many
SELECT sda.detection_type_id, COUNT(*) as alert_count
FROM "SafetyDetectionAlerts" sda
         JOIN "Drones" d ON sda.drone_id = d.drone_id
         JOIN "Customers" c ON d.customer_id = c.customer_id
WHERE c.customer_id = $1 AND sda.status != 'resolved'
GROUP BY sda.detection_type_id
ORDER BY sda.detection_type_id;

-- name: GetAllUnResolvedAlertsByDetectionType :many
SELECT sda.*
FROM "SafetyDetectionAlerts" sda
         JOIN "Drones" d ON sda.drone_id = d.drone_id
         JOIN "Customers" c ON d.customer_id = c.customer_id
WHERE c.customer_id = $1 AND sda.detection_type_id = $2;


-- name: GetAllOfficerIssues :many
SELECT i.*
FROM "Issues" i
         JOIN "SecurityOfficerIssues" soi ON i.issue_id = soi.issue_id
         JOIN "SecuritOfficers" so ON soi.officer_id = so.officer_id
         JOIN "Customers" c ON so.customer_id = c.customer_id
WHERE c.customer_id = $1
ORDER BY
    CASE
        WHEN i.status = 'New' THEN 1
        WHEN i.status = 'Pending' THEN 2
        WHEN i.status = 'Resolved' THEN 3
        ELSE 4
        END,
    i.created_at
    LIMIT $3 -- Number of issues per page
OFFSET $2; -- Page number (0-based)








