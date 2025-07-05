package posts

import (
	"context"
	"strconv"
	"strings"
	"time"
	"github.com/rs/zerolog/log"

	"github.com/dedimurphy/blog-api/internal/model/posts"
)

func (s *service) CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error {
	postHashtags := strings.Join(req.PostHashtags, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID: userId,
		PostTitle: req.PostTitle,
		PostContent: req.PostContent,
		PostHashtags: postHashtags,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userId, 10),
		UpdatedBy: strconv.FormatInt(userId, 10),
	}

	err := s.postsRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error create post to repository")
		return err
	}

	return nil
}