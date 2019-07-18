package repoprovince

import (
	"context"

	"github.com/fsetiawan29/healthy_food/domain/province"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/jmoiron/sqlx"
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
		return []models.Province{}, err
	}

	return result, nil
}
