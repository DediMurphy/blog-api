package posts

import (
	"context"

	"github.com/dedimurphy/fast-campus/internal/configs"
	"github.com/dedimurphy/fast-campus/internal/model/posts"
)

type postsRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context, id int64) (*posts.Post, error)

	CountLikedByPostID(ctx context.Context, postID int64) (int, error)

	GetCommentByPostID(ctx context.Context, postID int64) ([]posts.Comment, error)
}

type service struct {	
	cfg *configs.Config
	postsRepo postsRepository
}

 func NewService(cfg *configs.Config, postsRepo postsRepository) *service {
	return &service{
		cfg: cfg,
		postsRepo: postsRepo,
	}
 }
