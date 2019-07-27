package repoimage

import (
	"context"

	"github.com/fsetiawan29/healthy_food/domain/image"
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

func (m *mysqlImageRepository) GetImageByName(ctx context.Context, imageName string) (imageData models.Image, err error) {
	err = m.DB.GetContext(ctx, &imageData, image.QueryGetImageByName, imageName)
	if err != nil && !models.IsErrorNoRows(err) {
		logrus.Error(err)
		return models.Image{}, err
	}

	return imageData, nil
}

func (m *mysqlImageRepository) UpdateImage(ctx context.Context, imageData *models.Image) error {
	_, err := m.DB.NamedQuery(image.QueryUpdateImage, &imageData)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (m *mysqlImageRepository) GetImageByReferenceID(ctx context.Context, referenceID int64) (imageList []models.Image, err error) {
	err = m.DB.SelectContext(ctx, &imageList, image.QueryGetImageByReferenceID, referenceID)
	if err != nil && !models.IsErrorNoRows(err) {
		logrus.Error(err)
		return []models.Image{}, err
	}

	return imageList, nil
}

func (m *mysqlImageRepository) GetImageByReferenceIDRefType(ctx context.Context, referenceID int64, referenceType int) (imageList []models.Image, err error) {
	err = m.DB.SelectContext(ctx, &imageList, image.QueryGetImageByReferenceIDRefType, referenceID, referenceType)
	if err != nil && !models.IsErrorNoRows(err) {
		logrus.Error(err)
		return []models.Image{}, err
	}

	return imageList, nil
}
