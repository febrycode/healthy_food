package user

import (
	"context"

	"github.com/fsetiawan29/healthy_food/models"
)

// Repository represent the article's repository contract
type Repository interface {
	GetByEmail(ctx context.Context, email string) (models.User, error)
	GetByUserID(ctx context.Context, userID int64) (models.User, error)
	CreateUser(ctx context.Context, userData *models.User) error
	UpdateUser(ctx context.Context, userData *models.User) error
	GetListUser(ctx context.Context) (userList []models.User, err error)
}

type Usecase interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByUserID(ctx context.Context, userID int64) (models.User, error)
	CreateUser(ctx context.Context, userData *models.User) error
	UpdateUser(ctx context.Context, userData *models.User) error
	GetListUser(ctx context.Context) ([]models.User, error)
}
