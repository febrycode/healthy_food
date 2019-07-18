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

	foodList, err := uc.foodRepository.GetFood(ctx)
	if err != nil {
		return []models.FoodResponse{}, err
	}

	for _, foodData := range foodList {
		result = append(result, models.FoodResponse{
			Food: foodData,
		})
	}

	for i, resultData := range result {
		foodDetailList, err := uc.foodDetailRepository.GetFoodDetailByReferenceID(ctx, resultData.ID)
		if err != nil {
			return []models.FoodResponse{}, err
		}

		for _, foodDetailData := range foodDetailList {
			if foodDetailData.ReferenceType == 1 {
				result[i].Benefits = append(result[i].Benefits, models.Benefit{
					ID:            foodDetailData.ID,
					ReferenceType: foodDetailData.ReferenceType,
					ReferenceID:   resultData.ID,
					Description:   foodDetailData.Description,
				})
			}

			if foodDetailData.ReferenceType == 2 {
				result[i].Disadvantages = append(result[i].Disadvantages, models.Disadvantage{
					ID:            foodDetailData.ID,
					ReferenceType: foodDetailData.ReferenceType,
					ReferenceID:   resultData.ID,
					Description:   foodDetailData.Description,
				})
			}
		}

		imageList, err := uc.imageRepository.GetImageByReferenceID(ctx, resultData.ID)
		if err != nil {
			return []models.FoodResponse{}, err
		}

		for _, imageData := range imageList {
			result[i].Images = append(result[i].Images, imageData)
		}
	}

	return result, nil
}
