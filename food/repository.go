package food

import (
	"context"

	"github.com/febrycode/healthy_food/models"
)

type Repository interface {
	CreateFood(ctx context.Context, foodData *models.Food) (result int64, err error)
}
