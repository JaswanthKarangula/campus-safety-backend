-- name: DeleteDrone :one
DELETE FROM "Drones"
WHERE drone_id = $1 RETURNING *;

-- name: CreateNewCustomer :one
INSERT INTO "Customers" (customer_name, admin_id, resource_usage)
VALUES ($1,$2,$3)  RETURNING *;


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
SELECT so.*
FROM "SecuritOfficers" so
         JOIN "SecurityOfficerSchedule" sch ON so.officer_id = sch.officer_id
WHERE
        sch.day = EXTRACT(DOW FROM NOW())::varchar AND
    sch.start_time <= CURRENT_TIME AND
    sch.end_time >= CURRENT_TIME;



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










