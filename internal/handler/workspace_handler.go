package handler

import (
	"ide/internal/dto"
	"ide/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkspaceHandler struct {
	workspaceService *service.WorkspaceService
}

func newWorkspaceHandler(service *service.WorkspaceService) *WorkspaceHandler {
	return &WorkspaceHandler{workspaceService: service}
}

func (h *WorkspaceHandler) CreateWorkspace(ctx *gin.Context) {
	var payload dto.CreateWorkspacePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	accID, _ := ctx.Get("accountID")

	if err := h.workspaceService.CreateWorkspace(
		ctx,
		accID.(uuid.UUID),
		&payload,
	); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{"message": "successfully created workspace"},
	)
}
