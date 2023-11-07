// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type AlertHistory struct {
	HistoryID             int64          `json:"history_id"`
	AlertID               int64          `json:"alert_id"`
	Status                string         `json:"status"`
	ResolutionDescription sql.NullString `json:"resolution_description"`
	ResolvedBy            int64          `json:"resolved_by"`
	ResolutionDate        time.Time      `json:"resolution_date"`
}

type CustomerOfficerMapping struct {
	CustomerID int64 `json:"customer_id"`
	OfficerID  int64 `json:"officer_id"`
}

type DetectionType struct {
	DetectionID int64          `json:"detection_id"`
	Type        string         `json:"type"`
	Description sql.NullString `json:"description"`
}

type Drone struct {
	DroneID    int64        `json:"drone_id"`
	Model      string       `json:"model"`
	Status     string       `json:"status"`
	LastActive sql.NullTime `json:"last_active"`
}

type DroneLocation struct {
	LocationID int64     `json:"location_id"`
	DroneID    int64     `json:"drone_id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Timestamp  time.Time `json:"timestamp"`
}

type Notification struct {
	NotificationID int64     `json:"notification_id"`
	UserID         int64     `json:"user_id"`
	Message        string    `json:"message"`
	Timestamp      time.Time `json:"timestamp"`
	Status         string    `json:"status"`
	AlertID        int64     `json:"alert_id"`
}

type SafetyDetectionAlert struct {
	AlertID     int64          `json:"alert_id"`
	Timestamp   time.Time      `json:"timestamp"`
	FrameID     string         `json:"frame_id"`
	DroneID     int64          `json:"drone_id"`
	DetectionID int64          `json:"detection_id"`
	Description sql.NullString `json:"description"`
	Status      string         `json:"status"`
	ReviewedBy  int64          `json:"reviewed_by"`
}

type SecurityOfficerSchedule struct {
	ScheduleID int64     `json:"schedule_id"`
	OfficerID  int64     `json:"officer_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	// e.g., Mon,Tue,Wed
	Days string `json:"days"`
}

type StaffCustomerMapping struct {
	StaffID    int64 `json:"staff_id"`
	CustomerID int64 `json:"customer_id"`
}

type User struct {
	UserID         int64     `json:"user_id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Hashedpassword string    `json:"hashedpassword"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
}