package repository

import (
	"context"

	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/province"
	"github.com/jmoiron/sqlx"
	"github.com/tokopedia/tokopoints/errors"
)

type mysqlProvinceRepository struct {
	DB *sqlx.DB
}

func NewMysqlProvinceRepository(DB *sqlx.DB) province.Repository {
	return &mysqlProvinceRepository{
		DB: DB,
	}
}

func (m *mysqlProvinceRepository) GetAllProvince(ctx context.Context) (result []models.Province, err error) {
	err = m.DB.SelectContext(ctx, &result, province.QueryGetAllProvince)
	if err != nil {
		return []models.Province{}, errors.AddTrace(err)
	}

	return result, nil
}
