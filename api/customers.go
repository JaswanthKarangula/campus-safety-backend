package api

import (
	db "dronesaefty-backend/db/sqlc"
	"dronesaefty-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) SetUpCustomerRouter() {

	server.router.POST("/customer/addNewDrone", server.addNewDroneForCustomer)
	server.router.POST("/customer/createCustomer", server.createNewCustomer)
	server.router.GET("/customer/getAllActiveOfficers", server.getActiveOfficers)
	server.router.GET("/customer/getAllDrones", server.getAllDrones)
	server.router.GET("/customer/getAllOfficerIssues", server.getAllOfficerIssues)
	server.router.GET("/customer/getAllIssuesByCustomer", server.getAllIssuesByCustomer)

}

type GetAllIssuesRequest struct {
	CustomerID int64 `json:"customerid" binding:"required"`
	Limit      int32 `json:"limit" binding:"required"`
	Offset     int32 `json:"offset" binding:"required"`
}

// CreateTags		godoc
// @Summary			getAllIssuesByCustomer
// @Description 	returns all issues raised by Customer.
// @Param 			device body GetAllIssuesRequest true "returns all issues raised by customers  in Db"
// @Produce 		application/json
// @Tags 			customer
// @Success 		200 {object} string
// @Router			/customer/getAllIssuesByCustomer [get]
func (server *Server) getAllIssuesByCustomer(ctx *gin.Context) {

	var req GetAllIssuesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.GetAllIssuesByCustomerParams{
		CustomerID: req.CustomerID,
		Limit:      req.Limit,
		Offset:     req.Offset,
	}

	res, err := service.GetAllCustomerIssues(ctx, server.store, arg)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}

// CreateTags		godoc
// @Summary			getAllOfficerIssues
// @Description 	returns all issues raised by officers  of a Customer.
// @Param 			device body GetAllIssuesRequest true "returns all issues raised by officers  of a Customer in Db"
// @Produce 		application/json
// @Tags 			customer
// @Success 		200 {object} string
// @Router			/customer/getAllOfficerIssues [get]
func (server *Server) getAllOfficerIssues(ctx *gin.Context) {

	var req GetAllIssuesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.GetAllOfficerIssuesParams{
		CustomerID: req.CustomerID,
		Limit:      req.Limit,
		Offset:     req.Offset,
	}

	res, err := service.GetAllOfficerIssues(ctx, server.store, arg)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}

type GetAllDronesRequest struct {
	CustomerID int64 `json:"customerid" binding:"required"`
}

// CreateTags		godoc
// @Summary			getAllDrones
// @Description 	eturns all drones  of a Customer.
// @Param 			device body GetAllDronesRequest true "returns all drones  of a Customer  in Db"
// @Produce 		application/json
// @Tags 			customer
// @Success 		200 {object} string
// @Router			/customer/getAllDrones [get]
func (server *Server) getAllDrones(ctx *gin.Context) {

	var req GetAllDronesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := service.GetAllDrones(ctx, server.store, req.CustomerID)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}

type GetActiveOfficersRequest struct {
	CustomerID int64 `json:"customerid" binding:"required"`
}

// CreateTags		godoc
// @Summary			getActiveOfficers
// @Description 	returns all active security officers.
// @Param 			device body GetActiveOfficersRequest true "returns active security officers For Customer  in Db"
// @Produce 		application/json
// @Tags 			customer
// @Success 		200 {object} string
// @Router			/customer/getAllActiveOfficers [get]
func (server *Server) getActiveOfficers(ctx *gin.Context) {

	var req GetActiveOfficersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := service.GetActiveOfficers(ctx, server.store, req.CustomerID)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}

type CreateNewCustomerRequest struct {
	Username string `json:"username" binding:"required, alphanum"`
	Email    string `json:"email" binding:"required,min=6"`
	Password string `json:"password" binding:"required,email"`
}

// CreateTags		godoc
// @Summary			addNewDroneForCustomer
// @Description 	adds a New Drone For Customer  in Db.
// @Param 			device body AddNewDroneRequest true "adds a New Drone For Customer  in Db"
// @Produce 		application/json
// @Tags 			customer
// @Success 		200 {object} string
// @Router			/addNewDrone [post]
func (server *Server) createNewCustomer(ctx *gin.Context) {

	var req CreateNewCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateNewCustomerParams{
		Username:       req.Username,
		Hashedpassword: req.Password,
		Email:          req.Email,
	}

	res, err := service.CreateNewCustomer(ctx, server.store, arg)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}

type AddNewDroneRequest struct {
	Model      string `json:"model" binding:"required"`
	CustomerID int64  `json:"customer_id" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

// CreateTags		godoc
// @Summary			addNewDroneForCustomer
// @Description 	adds a New Drone For Customer  in Db.
// @Param 			device body AddNewDroneRequest true "adds a New Drone For Customer  in Db"
// @Produce 		application/json
// @Tags 			customer
// @Success 		200 {object} string
// @Router			/addNewDrone [post]
func (server *Server) addNewDroneForCustomer(ctx *gin.Context) {

	var req AddNewDroneRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.AddNewDroneParams{
		Model:      req.Model,
		Status:     req.Status,
		CustomerID: req.CustomerID,
	}

	res, err := service.AddNewDrone(ctx, server.store, arg)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}
