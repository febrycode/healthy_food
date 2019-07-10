package http

import (
	"context"
	"net/http"

	middlewareCustom "github.com/febrycode/healthy_food/middleware"
	"github.com/febrycode/healthy_food/province"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ProvinceHandler struct {
	provinceUsecase province.Usecase
}

func NewProvinceHandler(e *echo.Echo, provinceUsecase province.Usecase) {
	handler := &ProvinceHandler{
		provinceUsecase: provinceUsecase,
	}

	// Restricted group
	r := e.Group("/provinces")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &middlewareCustom.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("", handler.GetAllProvince)
}

func (p *ProvinceHandler) GetAllProvince(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := p.provinceUsecase.GetAllProvince(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
