package memberships

import (
	"context"

	"github.com/dedimurphy/blog-api/internal/middleware"
	"github.com/dedimurphy/blog-api/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string,string, error) 
	ValideteRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, memmembershipSvc membershipService) *Handler {
	return &Handler{
		Engine: api,
		membershipSvc: memmembershipSvc,
	}
}

// route
func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)

	routeRefresh := h.Group("memberships")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/refresh", h.Refresh)
}

