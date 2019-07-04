package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/febrycode/healthy_food/food_detail"
	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/user"
)

type mysqlFoodDetailRepository struct {
	DB *sqlx.DB
}

// NewMysqlFooDetailRepository will create an object that represent the user.Repository interface
func NewMysqlFooDetailRepository(DB *sqlx.DB) food_detail.Repository {
	return &mysqlFoodDetailRepository{DB}
}

func (m *mysqlFoodDetailRepository) CreateFoodDetail(ctx context.Context, foodDetailData *models.FoodDetail) error {
	_, err := m.DB.NamedQuery(user.QueryInsertUser, &foodDetailData)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
