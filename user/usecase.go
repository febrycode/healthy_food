package user

import (
	"context"

	"github.com/febrycode/healthy_food/models"
)

// Usecase represent the user's usecase
type Usecase interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	CreateUser(ctx context.Context, userData *models.User) error
}
