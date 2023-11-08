package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

func AddNewDrone(ctx *gin.Context, store db.Store, droneRequest db.AddNewDroneParams) (string, error) {

	return "Successfully added New Drone", nil
}
