package user

import (
	"context"

	"github.com/febrycode/healthy_food/models"
)

// Repository represent the article's repository contract
type Repository interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}
