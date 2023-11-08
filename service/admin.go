package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

func AddNewAdmin(ctx *gin.Context, store db.Store, arg db.CreateNewAdminParams) (db.User, error) {

	user_id, err := store.CreateNewAdmin(ctx, arg)

	if err != nil {
		return db.User{}, err
	}

	user, err := store.GetUser(ctx, user_id)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
