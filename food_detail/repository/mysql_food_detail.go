package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/febrycode/healthy_food/food_detail"
	"github.com/febrycode/healthy_food/models"
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
