package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

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
