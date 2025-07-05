package posts

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/dedimurphy/fast-campus/internal/model/posts"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postsRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id to database")
		return nil, err
	}

	likeCount, err := s.postsRepo.CountLikedByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like to database")
		return nil, err
	}

	comments, err := s.postsRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comment to database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments: comments,
	}, nil
}