package posts

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/dedimurphy/fast-campus/internal/model/posts"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := pageIndex * (pageIndex - 1)
	response, err := s.postsRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error get all post from database")
		return response, err
	}

	return response, nil
}