package api

import (
	db "dronesaefty-backend/db/sqlc"
	"dronesaefty-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) SetUpAdminRouter() {

	server.router.POST("/admin/createAdmin", server.createNewAdmin)
}

type CreateNewAdminRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// CreateTags		godoc
// @Summary			creates a new admin
// @Description 	adds a New admin in Db.
// @Param 			device body CreateNewAdminRequest true "adds a New Drone For Customer  in Db"
// @Produce 		application/json
// @Tags 			admin
// @Success 		200 {object} string
// @Router			/admin/createAdmin [post]
func (server *Server) createNewAdmin(ctx *gin.Context) {

	var req CreateNewAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateNewAdminParams{
		Username:       req.Username,
		Hashedpassword: req.Password,
		Email:          req.Email,
	}

	res, err := service.AddNewAdmin(ctx, server.store, arg)
	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}
