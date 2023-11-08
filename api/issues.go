package api

import (
	db "dronesaefty-backend/db/sqlc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) SetUpIssueRouter() {
	server.router.PUT("/issue/updateIssue", server.updateIssue)
}

type UpdateIssueRequest struct {
	Status   string `json:"status" binding:"required""`
	Comments string `json:"comments" binding:"required"`
	IssueID  int64  `json:"issue_id" binding:"required"`
}

// CreateTags		godoc
// @Summary			updateIssue
// @Description 	updates an issue by officer  in Db.
// @Param 			device body UpdateIssueRequest true "updates an  issue  in Db"
// @Produce 		application/json
// @Tags 			issue
// @Success 		200 {object} string
// @Router			/issue/updateIssue [put]
func (server *Server) updateIssue(ctx *gin.Context) {

	var req UpdateIssueRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println("In Request")

	arg := db.UpdateIssueStatusParams{
		Status:   req.Status,
		IssueID:  req.IssueID,
		Comments: req.Comments,
	}

	res, err := server.store.UpdateIssueStatus(ctx, arg)

	if err != nil {
		parseError(err, ctx)
	}

	ctx.JSON(http.StatusOK, res)
}
