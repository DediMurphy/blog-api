// use case
package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/dedimurphy/fast-campus/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipsRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username or email alreadry exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email     : req.Email,
		Username  : req.Username,
		Password  : string(pass),
		CreatedAt :now,
		UpdateAt  : now,
		CreatedBy : req.Email,
		UpdateBy  : req.Email,
	}

	return s.membershipsRepo.CreateUser(ctx, model)
}