package service

import (
	db "dronesaefty-backend/db/sqlc"
	"errors"
	"github.com/gin-gonic/gin"
)

func UpdateSchedules(ctx *gin.Context, store db.Store, schedule []db.UpdateScheduleParams) (string, error) {
	if len(schedule) == 0 {
		return "", errors.New("Schedule is empty")
	}

	_, err := store.DeleteOfficerSchedule(ctx, schedule[0].OfficerID)

	if err != nil {
		return "", err
	}

	for _, params := range schedule {
		_, err := store.UpdateSchedule(ctx, params)
		if err != nil {
			return "", err
		}
	}

	return "Successfully Updated Schedules", nil
}

func CreateNewOfficerIssue(ctx *gin.Context, store db.Store, req db.CreateSecurityOfficerIssueParams) (string, error) {
	_, err := store.CreateSecurityOfficerIssue(ctx, req)
	if err != nil {
		return "", err
	}
	return "Successfully Raised Issue", nil
}

func CreateNewSecurityOfficer(ctx *gin.Context, store db.Store, arg db.CreateNewSecurityOfficerParams) (db.User, error) {

	securityOfficer, err := store.CreateNewSecurityOfficer(ctx, arg)

	if err != nil {
		return db.User{}, err
	}

	user, err := store.GetUser(ctx, securityOfficer.OfficerID)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
