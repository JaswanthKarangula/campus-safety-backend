package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

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
