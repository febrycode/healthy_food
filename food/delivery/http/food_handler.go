package http

import (
	"context"

	"github.com/febrycode/healthy_food/food"
	"github.com/febrycode/healthy_food/models"
	"github.com/labstack/echo"
)

type FoodHandler struct {
	foodUsecase food.Usecase
}

func NewFoodHandler(e *echo.Echo, foodUsecase food.Usecase) {
	handler := &FoodHandler{
		foodUsecase: foodUsecase,
	}

	e.POST("/food", handler.CreateFood)
}

func (f *FoodHandler) CreateFood(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	foodParam := &models.FoodRequest{}
	if err = c.Bind(foodParam); err != nil {
		return err
	}

	err = f.foodUsecase.CreateFood(ctx, foodParam)
	if err != nil {
		return err
	}

	return nil
}
