package handler

import (
	"errors"
	"kroff/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	UsernameCtx         = "username"
	RoleCtx             = "role"
)

var (
	errUnauthorized = errors.New("unauthorized")
)

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,Access-Control-Request-Method, Access-Control-Request-Headers")
		ctx.Header("Access-Control-Max-Age", "3600")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH, HEAD")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}

func (h *Handler) adminIdentity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(AuthorizationHeader)
		if header == "" {
			response.ErrorResponse(ctx, http.StatusUnauthorized, errUnauthorized)
			ctx.Abort()
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			response.ErrorResponse(ctx, http.StatusUnauthorized, errUnauthorized)
			ctx.Abort()
			return
		}

		if len(headerParts[1]) == 0 {
			response.ErrorResponse(ctx, http.StatusUnauthorized, errUnauthorized)
			ctx.Abort()
			return
		}

		claims, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			response.ErrorResponse(ctx, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		if claims.Role != "admin" {
			response.ErrorResponse(ctx, http.StatusUnauthorized, errUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Set(UsernameCtx, claims.Username)
		ctx.Set(RoleCtx, claims.Role)
		ctx.Next()
	}
}
