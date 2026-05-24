package handler

import (
	"ide/internal/dto"
	"ide/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService *service.AuthService
}

func newAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	var payload dto.SignInPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	token, err := h.authService.SignIn(ctx.Request.Context(), &payload)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.SetCookie("token", token, 0, "/", "", true, true)
	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "successfully signed in"},
	)
}

func (h *AuthHandler) CreateProfile(ctx *gin.Context) {
	var payload dto.CreateProfilePayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	accID, _ := ctx.Get("accountID")

	if err := h.authService.CreateProfile(ctx, accID.(uuid.UUID), &payload); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "successfully created profile"},
	)
}

func (h *AuthHandler) GetProfile(ctx *gin.Context) {
	accID, _ := ctx.Get("accountID")

	profile, err := h.authService.GetProfile(ctx, accID.(uuid.UUID))
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		profile,
	)
}
