package usecase

import (
	"context"
	"time"

	"github.com/febrycode/healthy_food/image"
	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/util"
)

type Usecase struct {
	imageRepository image.Repository
	contextTimeout  time.Duration
}

func NewImageUsecase(i image.Repository, timeout time.Duration) image.Usecase {
	return &Usecase{
		imageRepository: i,
		contextTimeout:  timeout,
	}
}

func (uc *Usecase) CreateImage(ctx context.Context, imageData *models.Image) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	imageData.CreatedAt = util.GetTimeNow()

	err := uc.imageRepository.CreateImage(ctx, imageData)
	if err != nil {
		return err
	}

	return nil
}
