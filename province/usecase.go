package province

import (
	"context"

	"github.com/fsetiawan29/healthy_food/models"
)

type Usecase interface {
	GetAllProvince(ctx context.Context) (result []models.Province, err error)
}
