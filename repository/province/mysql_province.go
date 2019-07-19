package repoprovince

import (
	"context"

	"github.com/fsetiawan29/healthy_food/domain/province"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (m *mysqlProvinceRepository) GetProvinceByID(ctx context.Context, provinceID int64) (provinceData models.Province, err error) {
	err = m.DB.GetContext(ctx, &provinceData, province.QueryGetProvinceByID, provinceID)
	if err != nil && !models.IsErrorNoRows(err) {
		logrus.Error(err)
		return models.Province{}, err
	}

	return provinceData, nil
}
