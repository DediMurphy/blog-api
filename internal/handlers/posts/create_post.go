package posts

import (
	"net/http"

	"github.com/dedimurphy/blog-api/internal/model/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := c.GetInt64("userID")
	err := h.postSvc.CreatePost(ctx, userId, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusAccepted)
}