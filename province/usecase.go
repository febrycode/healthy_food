package province

import (
	"context"

	"github.com/febrycode/healthy_food/models"
)

type Usecase interface {
	GetAllProvince(ctx context.Context) (result []models.Province, err error)
}
