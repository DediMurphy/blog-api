package posts

import (
	"context"

	"github.com/dedimurphy/blog-api/internal/middleware"
	"github.com/dedimurphy/blog-api/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID int64, userID int64, req posts.CreateCommentsRequest) error
	UpsertUpdateUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error

	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error)
}

type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine: api,
		postSvc: postSvc,
	}
}

// route
func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	//middleware
	route.Use(middleware.AuthMiddleware())
	route.POST("/create", h.CreatePost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/user_activity/:postID", h.UpsertUserActivity)
	route.GET("/", h.GetAllPost)
	route.GET("/:postID", h.GetPostByID)
}

