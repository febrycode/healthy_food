package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/fsetiawan29/healthy_food/food"
	"github.com/fsetiawan29/healthy_food/food_detail"
	"github.com/fsetiawan29/healthy_food/models"
)

type Usecase struct {
	foodRepository       food.Repository
	foodDetailRepository food_detail.Repository
	contextTimeout       time.Duration
}

func NewFoodUsecase(f food.Repository, fd food_detail.Repository, timeout time.Duration) food.Usecase {
	return &Usecase{
		foodRepository:       f,
		foodDetailRepository: fd,
		contextTimeout:       timeout,
	}
}

func (uc *Usecase) CreateFood(ctx context.Context, foodParam *models.FoodRequest) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	foodID, err := uc.foodRepository.CreateFood(ctx, food.BuilderFoodParamToFood(foodParam))
	if err != nil {
		fmt.Println(err)
		return err
	}

	foodDetailList := food.BuilderFoodParamToFoodDetail(foodID, foodParam)
	for _, foodDetailData := range foodDetailList {
		err = uc.foodDetailRepository.CreateFoodDetail(ctx, foodDetailData)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *Usecase) GetFood(ctx context.Context) (result []models.FoodResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	result, err = uc.foodRepository.GetFood(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
