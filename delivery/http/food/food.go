package httpfood

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fsetiawan29/healthy_food/domain/food"
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

	e.GET("/public/food", handler.GetFood)

	// Restricted group
	r := e.Group("/food")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &middlewareCustom.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.POST("", handler.CreateFood)
	r.GET("", handler.GetFood)

}

func (f *FoodHandler) CreateFood(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	foodParam := &models.FoodRequest{}
	if err = c.Bind(foodParam); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewareCustom.JwtCustomClaims)

	foodParam.UserID = claims.UserID
	err = f.foodUsecase.CreateFood(ctx, foodParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	return c.JSON(http.StatusCreated, models.ResponseJSON(http.StatusCreated, "Food created successfully"))
}

func (f *FoodHandler) GetFood(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var result []models.FoodResponse

	title := c.QueryParam("title")
	if title == "" {
		result, err = f.foodUsecase.GetFood(ctx)
		if err != nil {
			return err
		}
	} else {
		result, err = f.foodUsecase.GetFoodByTitle(ctx, title)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, result)
}
