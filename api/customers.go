package api

import (
	"dronesaefty-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	res, err := service.AddNewDrone(ctx, server.store, req)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}
