package food

import (
	"context"

	"github.com/fsetiawan29/healthy_food/models"
)

type Repository interface {
	CreateFood(ctx context.Context, foodData *models.Food) (result int64, err error)
	GetFood(ctx context.Context) (foodList []models.Food, err error)
}

type Usecase interface {
	CreateFood(ctx context.Context, foodParam *models.FoodRequest) error
	GetFood(ctx context.Context) (result []models.FoodResponse, err error)
}
