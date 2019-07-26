package repofood

import (
	"context"

	"github.com/fsetiawan29/healthy_food/domain/food"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/jmoiron/sqlx"
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

func (m *mysqlUserRepository) GetFood(ctx context.Context) (foodList []models.Food, err error) {
	err = m.DB.SelectContext(ctx, &foodList, food.QueryGetListFood)
	if err != nil {
		return []models.Food{}, err
	}

	return foodList, nil
}
