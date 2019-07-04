package usecase

import (
	"context"
	"time"

	"github.com/febrycode/healthy_food/food"
	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/util"
)

type Usecase struct {
	foodRepository food.Repository
	contextTimeout time.Duration
}

func NewFoodUsecase(f food.Repository, timeout time.Duration) food.Usecase {
	return &Usecase{
		foodRepository: f,
		contextTimeout: timeout,
	}
}

func (uc *Usecase) CreateFood(ctx context.Context, foodData *models.Food) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	foodData.CreatedAt = util.GetTimeNow()

	err := uc.foodRepository.CreateFood(ctx, foodData)
	if err != nil {
		return err
	}

	return nil
}
