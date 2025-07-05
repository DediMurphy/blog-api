package memberships

import (
	"context"
	"time"

	"github.com/dedimurphy/fast-campus/internal/configs"
	"github.com/dedimurphy/fast-campus/internal/model/memberships"
)

type membershipsRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error

	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
}

type service struct {	
	cfg *configs.Config
	membershipsRepo membershipsRepository
}

 func NewService(cfg *configs.Config, membershipRepo membershipsRepository) *service {
	return &service{
		cfg: cfg,
		membershipsRepo: membershipRepo,
	}
 }
