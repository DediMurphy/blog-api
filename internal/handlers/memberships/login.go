package memberships

import (
	"net/http"

	"github.com/dedimurphy/blog-api/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := h.membershipSvc.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),		
		})
	}

	response := memberships.LoginResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, response)
}