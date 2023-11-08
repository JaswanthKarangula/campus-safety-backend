package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

func GetActiveOfficers(ctx *gin.Context, store db.Store, customerID int64) ([]db.User, error) {

	activeCustomers, err := store.GetAllActiveSecurityOfficers(ctx, customerID)

	if err != nil {
		return []db.User{}, err
	}

	return activeCustomers, nil
}

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
