package usecase

import (
	"context"
	"time"

	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/user"
)

type userUsecase struct {
	userRepository user.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(u user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepository: u,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return res, nil
}
