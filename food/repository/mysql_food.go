package repository

import (
	"context"

	"github.com/febrycode/healthy_food/food"
	"github.com/febrycode/healthy_food/models"
	"github.com/jmoiron/sqlx"
	"github.com/tokopedia/tokopoints/errors"
)

type mysqlUserRepository struct {
	DB *sqlx.DB
}

func NewMysqlFoodRepository(DB *sqlx.DB) food.Repository {
	return &mysqlUserRepository{
		DB: DB,
	}
}

func (m *mysqlUserRepository) CreateFood(ctx context.Context, foodData *models.Food) (result int64, err error) {
	rows, err := m.DB.NamedQuery(food.QueryInsertFood, foodData)
	if err != nil {
		return 0, errors.AddTrace(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return 0, errors.AddTrace(err)
		}
	}

	return 0, nil
}
