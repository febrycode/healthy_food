package repository

import (
	"context"

	"github.com/fsetiawan29/healthy_food/food"
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
