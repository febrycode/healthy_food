package repository

import (
	"context"

	"github.com/fsetiawan29/healthy_food/image"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type mysqlImageRepository struct {
	DB *sqlx.DB
}

func NewMysqlImageRepository(DB *sqlx.DB) image.Repository {
	return &mysqlImageRepository{
		DB: DB,
	}
}

func (m *mysqlImageRepository) CreateImage(ctx context.Context, imageData *models.Image) error {
	_, err := m.DB.NamedQuery(image.QueryInsertImage, imageData)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
