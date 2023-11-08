package service

import (
	"dronesaefty-backend/api"
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

func AddNewDrone(ctx *gin.Context, store db.Store, droneRequest api.AddNewDroneRequest) (string, error) {

	return "Successfully added New Drone", nil
}
