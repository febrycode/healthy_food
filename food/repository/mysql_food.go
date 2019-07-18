package repository

import (
	"context"

	"github.com/fsetiawan29/healthy_food/food"
	"github.com/fsetiawan29/healthy_food/food_detail"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type mysqlUserRepository struct {
	DB *sqlx.DB
}

func NewMysqlFoodRepository(DB *sqlx.DB) food.Repository {
	return &mysqlUserRepository{
		DB: DB,
	}
}

func (m *mysqlUserRepository) CreateFood(ctx context.Context, foodData *models.Food) (int64, error) {
	res, err := m.DB.NamedExec(food.QueryInsertFood, foodData)
	if err != nil {
		return 0, err
	}

	result, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (m *mysqlUserRepository) GetFood(ctx context.Context) ([]models.FoodResponse, error) {
	var foodList []models.Food
	err := m.DB.SelectContext(ctx, &foodList, food.QueryGetListFood)
	if err != nil {
		return []models.FoodResponse{}, err
	}

	var result []models.FoodResponse
	for _, foodData := range foodList {
		result = append(result, models.FoodResponse{
			Food: foodData,
		})
	}

	for i, resultData := range result {
		var foodDetailList []models.FoodDetail
		err = m.DB.SelectContext(ctx, &foodDetailList, food_detail.QueryGetFoodDetailByFoodID, resultData.ID)
		if err != nil && !models.IsErrorNoRows(err) {
			logrus.Error(err)
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

	}

	return result, nil
}
