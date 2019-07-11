package http

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fsetiawan29/healthy_food/food"
	middlewareCustom "github.com/fsetiawan29/healthy_food/middleware"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type FoodHandler struct {
	foodUsecase food.Usecase
}

func NewFoodHandler(e *echo.Echo, foodUsecase food.Usecase) {
	handler := &FoodHandler{
		foodUsecase: foodUsecase,
	}

	// Restricted group
	r := e.Group("/food")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &middlewareCustom.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.POST("", handler.CreateFood)

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

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewareCustom.JwtCustomClaims)

	foodParam.UserID = claims.UserID
	err = f.foodUsecase.CreateFood(ctx, foodParam)
	if err != nil {
		return err
	}

	return nil
}
