package image

import (
	"context"

	"github.com/fsetiawan29/healthy_food/models"
)

type Repository interface {
	CreateImage(ctx context.Context, imageData *models.Image) error
}
