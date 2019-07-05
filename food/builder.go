package food

import (
	"github.com/febrycode/healthy_food/common/enums/fooddetailreferencetype"
	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/util"
)

func BuilderFoodParamToFood(foodParam *models.FoodRequest) *models.Food {
	return &models.Food{
		ProvinceID: foodParam.ProvinceID,
		Title:      foodParam.Title,
		CreatedAt:  util.GetTimeNow(),
	}
}

func BuilderFoodParamToFoodDetail(referenceID int64, foodParam *models.FoodRequest) []*models.FoodDetail {
	foodDetailList := make([]*models.FoodDetail, 0)

	for _, benefitData := range foodParam.Benefit {
		foodDetailList = append(foodDetailList, &models.FoodDetail{
			ReferenceType: int(fooddetailreferencetype.BENEFIT),
			ReferenceID:   referenceID,
			Description:   benefitData,
			CreatedAt:     util.GetTimeNow(),
		})
	}

	for _, disadvantageData := range foodParam.Disadvantage {
		foodDetailList = append(foodDetailList, &models.FoodDetail{
			ReferenceType: int(fooddetailreferencetype.DISADVANTAGE),
			ReferenceID:   referenceID,
			Description:   disadvantageData,
			CreatedAt:     util.GetTimeNow(),
		})
	}

	return foodDetailList
}
