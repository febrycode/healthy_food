package food

import (
	"context"

	"github.com/fsetiawan29/healthy_food/models"
)

type Repository interface {
	CreateFood(ctx context.Context, foodData *models.Food) (result int64, err error)
	GetFood(ctx context.Context) (foodList []models.Food, err error)
}
