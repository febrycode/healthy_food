package repofooddetail

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/fsetiawan29/healthy_food/domain/food_detail"
	"github.com/fsetiawan29/healthy_food/models"
)

type mysqlFoodDetailRepository struct {
	DB *sqlx.DB
}

// NewMysqlFooDetailRepository will create an object that represent the user.Repository interface
func NewMysqlFoodDetailRepository(DB *sqlx.DB) food_detail.Repository {
	return &mysqlFoodDetailRepository{DB}
}

func (m *mysqlFoodDetailRepository) CreateFoodDetail(ctx context.Context, foodDetailData *models.FoodDetail) error {
	_, err := m.DB.NamedQuery(food_detail.QueryInsertFoodDetail, foodDetailData)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (m *mysqlFoodDetailRepository) GetFoodDetailByReferenceID(
	ctx context.Context,
	referenceID int64,
) (foodDetailList []models.FoodDetail, err error) {

	err = m.DB.SelectContext(ctx, &foodDetailList, food_detail.QueryGetFoodDetailByFoodID, referenceID)
	if err != nil && !models.IsErrorNoRows(err) {
		logrus.Error(err)
		return []models.FoodDetail{}, err
	}

	return foodDetailList, nil
}
