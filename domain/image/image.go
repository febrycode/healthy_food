package image

import (
	"context"

	"github.com/fsetiawan29/healthy_food/models"
)

type Repository interface {
	CreateImage(ctx context.Context, imageData *models.Image) error
	GetImageByName(ctx context.Context, imageName string) (imageData models.Image, err error)
	UpdateImage(ctx context.Context, imageData *models.Image) error
	GetImageByReferenceID(ctx context.Context, referenceID int64) (imageList []models.Image, err error)
	GetImageByReferenceIDRefType(ctx context.Context, referenceID int64, referenceType int) (imageList []models.Image, err error)
}

type Usecase interface {
	CreateImage(ctx context.Context, imageData *models.Image) error
}
