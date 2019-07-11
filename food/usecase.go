package food

import (
	"context"

	"github.com/fsetiawan29/healthy_food/models"
)

type Usecase interface {
	CreateFood(ctx context.Context, foodParam *models.FoodRequest) error
}
