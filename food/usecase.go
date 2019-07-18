package food

import (
	"context"

	"github.com/fsetiawan29/healthy_food/models"
)

type Usecase interface {
	CreateFood(ctx context.Context, foodParam *models.FoodRequest) error
	GetFood(ctx context.Context) (result []models.FoodResponse, err error)
}
