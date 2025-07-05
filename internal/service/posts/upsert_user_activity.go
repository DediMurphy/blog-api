package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/dedimurphy/fast-campus/internal/model/posts"
)

func (s *service) UpsertUpdateUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()

	model := posts.UserActivityModel {
		PostID: postID,
		UserID: userID,
		IsLiked: request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	userActivity, err := s.postsRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error liked post")
		return err
	}

	if userActivity == nil {
		if !request.IsLiked {
			return errors.New("anda belum pernah like sebelumnya")
		}
		err = s.postsRepo.CreateUserActivity(ctx, model)	
	} else {
		err = s.postsRepo.UpdateUserActivity(ctx, model)
	}
	if err != nil {
		log.Error().Err(err).Msg("error create or update user activity to database")
		return err
	}
	return nil
}