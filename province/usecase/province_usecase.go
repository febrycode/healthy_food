package usecase

import (
	"context"
	"time"

	"github.com/fsetiawan29/healthy_food/models"
	"github.com/fsetiawan29/healthy_food/province"
)

type Usecase struct {
	provinceRepository province.Repository
	contextTimeout     time.Duration
}

func NewProvinceUsecase(p province.Repository, timeout time.Duration) province.Usecase {
	return &Usecase{
		provinceRepository: p,
		contextTimeout:     timeout,
	}
}

func (uc *Usecase) GetAllProvince(ctx context.Context) (result []models.Province, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	result, err = uc.provinceRepository.GetAllProvince(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
