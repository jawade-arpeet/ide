package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "token not found"},
			)
			return
		}

		accID, err := uuid.Parse(token)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "invalid token"},
			)
			return
		}

		ctx.Set("accountID", accID)
		ctx.Next()
	}
}
