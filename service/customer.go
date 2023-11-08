package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

func CreateNewCustomer(ctx *gin.Context, store db.Store, arg db.CreateNewCustomerParams) (db.User, error) {

	customer, err := store.CreateNewCustomer(ctx, arg)

	if err != nil {
		return db.User{}, err
	}

	user, err := store.GetUser(ctx, customer.CustomerID)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func AddNewDrone(ctx *gin.Context, store db.Store, droneRequest db.AddNewDroneParams) (string, error) {

	return "Successfully added New Drone", nil
}
