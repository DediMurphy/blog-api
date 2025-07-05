package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/dedimurphy/blog-api/internal/model/memberships"
	"github.com/dedimurphy/blog-api/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValideteRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipsRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	// means the token in database is not mathchend
	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token has expired")
	}

	user, err := s.membershipsRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not exist")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJwt)
	if err != nil {
		return "", err
	}
	return token, nil
}