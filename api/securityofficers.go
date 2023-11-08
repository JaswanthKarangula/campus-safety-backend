package api

import (
	db "dronesaefty-backend/db/sqlc"
	"dronesaefty-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) SetUpSecurityOfficerRouter() {

	server.router.POST("/officer/createOfficer", server.createNewAdmin)
	server.router.POST("/officer/raiseOfficerIssue", server.createNewOfficerIssue)
}

type CreateNewOfficerIssueRequest struct {
	Description       string `json:"description" binding:"required"`
	SecurityOfficerID int64  `json:"officer_id" binding:"required"`
}

// CreateTags		godoc
// @Summary			createNewOfficerIssue
// @Description 	raises a new issue by officer  in Db.
// @Param 			device body CreateNewOfficerIssueRequest true "raises a new issue by officer in Db"
// @Produce 		application/json
// @Tags 			customer
// @Success 		200 {object} string
// @Router			/officer/raiseOfficerIssue [post]
func (server *Server) createNewOfficerIssue(ctx *gin.Context) {

	var req CreateNewOfficerIssueRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateSecurityOfficerIssueParams{
		OfficerID:   req.SecurityOfficerID,
		Description: req.Description,
	}

	res, err := service.CreateNewOfficerIssue(ctx, server.store, arg)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}

type CreateNewSecurityOfficerRequest struct {
	Username   string `json:"username" binding:"required, alphanum"`
	Email      string `json:"email" binding:"required,min=6"`
	Password   string `json:"password" binding:"required,email"`
	CustomerID int64  `json:"customerid" binding:"required"`
}

// CreateTags		godoc
// @Summary			createNewSecurityOfficer
// @Description 	adds a New Security officer For Customer  in Db.
// @Param 			device body CreateNewSecurityOfficerRequest true "adds a New officer For Customer  in Db"
// @Produce 		application/json
// @Tags 			officer
// @Success 		200 {object} string
// @Router			/officer/createOfficer [post]
func (server *Server) createNewSecurityOfficer(ctx *gin.Context) {

	var req CreateNewSecurityOfficerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateNewSecurityOfficerParams{
		Username:       req.Username,
		Hashedpassword: req.Password,
		Email:          req.Email,
		CustomerID:     req.CustomerID,
	}

	res, err := service.CreateNewSecurityOfficer(ctx, server.store, arg)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}
