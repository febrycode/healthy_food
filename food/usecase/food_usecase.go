package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/fsetiawan29/healthy_food/food"
	"github.com/fsetiawan29/healthy_food/food_detail"
	"github.com/fsetiawan29/healthy_food/image"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/fsetiawan29/healthy_food/util"
)

type Usecase struct {
	foodRepository       food.Repository
	foodDetailRepository food_detail.Repository
	contextTimeout       time.Duration
	imageRepository      image.Repository
}

func NewFoodUsecase(f food.Repository, fd food_detail.Repository, timeout time.Duration, i image.Repository) food.Usecase {
	return &Usecase{
		foodRepository:       f,
		foodDetailRepository: fd,
		contextTimeout:       timeout,
		imageRepository:      i,
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

	for _, foodImage := range foodParam.Images {
		imageData, err := uc.imageRepository.GetImageByName(ctx, foodImage)
		if err != nil {
			return err
		}

		imageData.ReferenceID = foodID
		imageData.ReferenceType = 2
		imageData.UpdatedAt = util.GetTimeNow()
		err = uc.imageRepository.UpdateImage(ctx, &imageData)
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
