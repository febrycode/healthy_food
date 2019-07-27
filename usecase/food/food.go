package ucfood

import (
	"context"
	"fmt"
	"time"

	"github.com/fsetiawan29/healthy_food/domain/food"
	"github.com/fsetiawan29/healthy_food/domain/food_detail"
	"github.com/fsetiawan29/healthy_food/domain/image"
	"github.com/fsetiawan29/healthy_food/domain/province"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/fsetiawan29/healthy_food/util"
)

type Usecase struct {
	foodRepository       food.Repository
	foodDetailRepository food_detail.Repository
	contextTimeout       time.Duration
	imageRepository      image.Repository
	provinceRepository   province.Repository
}

func NewFoodUsecase(
	f food.Repository,
	fd food_detail.Repository,
	timeout time.Duration,
	i image.Repository,
	p province.Repository,
) food.Usecase {

	return &Usecase{
		foodRepository:       f,
		foodDetailRepository: fd,
		contextTimeout:       timeout,
		imageRepository:      i,
		provinceRepository:   p,
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
		provinceData, err := uc.provinceRepository.GetProvinceByID(ctx, resultData.ProvinceID)
		if err != nil {
			return []models.FoodResponse{}, err
		}

		if provinceData.ID > 0 {
			result[i].ProvinceName = provinceData.Name
		}

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

	for i, resultData := range result {
		if len(resultData.Benefits) == 0 {
			result[i].Benefits = []models.Benefit{}
		}

		if len(resultData.Disadvantages) == 0 {
			result[i].Disadvantages = []models.Disadvantage{}
		}

		if len(resultData.Images) == 0 {
			result[i].Images = []models.Image{}
		}
	}

	return result, nil
}

func (uc *Usecase) GetFoodByTitle(ctx context.Context, title string) (result []models.FoodResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	foodList, err := uc.foodRepository.GetFoodByTitle(ctx, title)
	if err != nil {
		return []models.FoodResponse{}, err
	}

	for _, foodData := range foodList {
		result = append(result, models.FoodResponse{
			Food: foodData,
		})
	}

	for i, resultData := range result {
		provinceData, err := uc.provinceRepository.GetProvinceByID(ctx, resultData.ProvinceID)
		if err != nil {
			return []models.FoodResponse{}, err
		}

		if provinceData.ID > 0 {
			result[i].ProvinceName = provinceData.Name
		}

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
