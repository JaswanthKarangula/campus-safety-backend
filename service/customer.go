package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

func CreateNewCustomerIssue(ctx *gin.Context, store db.Store, req db.CreateCustomerIssueParams) (string, error) {
	_, err := store.CreateCustomerIssue(ctx, req)
	if err != nil {
		return "", err
	}
	return "Successfully Raised Issue", nil
}

func GetAllCustomerIssues(ctx *gin.Context, store db.Store, req db.GetAllIssuesByCustomerParams) ([]db.Issue, error) {

	activeCustomers, err := store.GetAllIssuesByCustomer(ctx, req)

	if err != nil {
		return []db.Issue{}, err
	}

	return activeCustomers, nil
}

func GetAllOfficerIssues(ctx *gin.Context, store db.Store, req db.GetAllOfficerIssuesParams) ([]db.Issue, error) {

	activeCustomers, err := store.GetAllOfficerIssues(ctx, req)

	if err != nil {
		return []db.Issue{}, err
	}

	return activeCustomers, nil
}

func GetAllDrones(ctx *gin.Context, store db.Store, customerID int64) ([]db.Drone, error) {

	activeCustomers, err := store.GetAllDronesByCustomer(ctx, customerID)

	if err != nil {
		return []db.Drone{}, err
	}

	return activeCustomers, nil
}

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
