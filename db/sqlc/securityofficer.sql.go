// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: securityofficer.sql

package db

import (
	"context"
	"time"
)

const createNewSecurityOfficer = `-- name: CreateNewSecurityOfficer :one
WITH new_security_officer AS (
INSERT INTO "Users" (username, email, hashedpassword, role)
VALUES ($1, $2, $3, 'Security Officer')
    RETURNING user_id, username, email, hashedpassword, role, created_at
    )
INSERT INTO "SecuritOfficers" (officer_id, customer_id)
SELECT user_id, $4
FROM new_security_officer
RETURNING officer_id, customer_id
`

type CreateNewSecurityOfficerParams struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Hashedpassword string `json:"hashedpassword"`
	CustomerID     int64  `json:"customer_id"`
}

func (q *Queries) CreateNewSecurityOfficer(ctx context.Context, arg CreateNewSecurityOfficerParams) (SecuritOfficer, error) {
	row := q.db.QueryRowContext(ctx, createNewSecurityOfficer,
		arg.Username,
		arg.Email,
		arg.Hashedpassword,
		arg.CustomerID,
	)
	var i SecuritOfficer
	err := row.Scan(&i.OfficerID, &i.CustomerID)
	return i, err
}

const createSecurityOfficerIssue = `-- name: CreateSecurityOfficerIssue :one
WITH new_issue AS (
INSERT INTO "Issues" (description, status, comments)
VALUES ($1, 'New', '')
    RETURNING issue_id
    )
INSERT INTO "SecurityOfficerIssues" (officer_id, issue_id)
SELECT $2, issue_id
FROM new_issue
RETURNING customer_id, officer_id, issue_id
`

type CreateSecurityOfficerIssueParams struct {
	Description string `json:"description"`
	OfficerID   int64  `json:"officer_id"`
}

func (q *Queries) CreateSecurityOfficerIssue(ctx context.Context, arg CreateSecurityOfficerIssueParams) (SecurityOfficerIssue, error) {
	row := q.db.QueryRowContext(ctx, createSecurityOfficerIssue, arg.Description, arg.OfficerID)
	var i SecurityOfficerIssue
	err := row.Scan(&i.CustomerID, &i.OfficerID, &i.IssueID)
	return i, err
}

const deleteOfficerSchedule = `-- name: DeleteOfficerSchedule :many
DELETE FROM "SecurityOfficerSchedule"
WHERE officer_id = $1
RETURNING schedule_id, officer_id, start_time, end_time, day
`

func (q *Queries) DeleteOfficerSchedule(ctx context.Context, officerID int64) ([]SecurityOfficerSchedule, error) {
	rows, err := q.db.QueryContext(ctx, deleteOfficerSchedule, officerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SecurityOfficerSchedule{}
	for rows.Next() {
		var i SecurityOfficerSchedule
		if err := rows.Scan(
			&i.ScheduleID,
			&i.OfficerID,
			&i.StartTime,
			&i.EndTime,
			&i.Day,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllIssuesByAllSecurityOfficers = `-- name: GetAllIssuesByAllSecurityOfficers :many
SELECT i.issue_id, i.description, i.status, i.comments, i.created_at
FROM "Issues" i
         JOIN "SecurityOfficerIssues" soi ON i.issue_id = soi.issue_id
         JOIN "SecuritOfficers" so ON soi.officer_id = so.officer_id
         JOIN "Customers" c ON so.customer_id = c.customer_id
WHERE c.customer_id = $1
`

func (q *Queries) GetAllIssuesByAllSecurityOfficers(ctx context.Context, customerID int64) ([]Issue, error) {
	rows, err := q.db.QueryContext(ctx, getAllIssuesByAllSecurityOfficers, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Issue{}
	for rows.Next() {
		var i Issue
		if err := rows.Scan(
			&i.IssueID,
			&i.Description,
			&i.Status,
			&i.Comments,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSchedule = `-- name: UpdateSchedule :one
UPDATE "SecurityOfficerSchedule"
SET start_time = $2,
    end_time = $3,
    day = $4
WHERE officer_id = $1
RETURNING schedule_id, officer_id, start_time, end_time, day
`

type UpdateScheduleParams struct {
	OfficerID int64     `json:"officer_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Day       string    `json:"day"`
}

func (q *Queries) UpdateSchedule(ctx context.Context, arg UpdateScheduleParams) (SecurityOfficerSchedule, error) {
	row := q.db.QueryRowContext(ctx, updateSchedule,
		arg.OfficerID,
		arg.StartTime,
		arg.EndTime,
		arg.Day,
	)
	var i SecurityOfficerSchedule
	err := row.Scan(
		&i.ScheduleID,
		&i.OfficerID,
		&i.StartTime,
		&i.EndTime,
		&i.Day,
	)
	return i, err
}
