package http

import (
	"context"
	"net/http"

	"github.com/febrycode/healthy_food/province"
	"github.com/labstack/echo"
)

type ProvinceHandler struct {
	provinceUsecase province.Usecase
}

func NewProvinceHandler(e *echo.Echo, provinceUsecase province.Usecase) {
	handler := &ProvinceHandler{
		provinceUsecase: provinceUsecase,
	}

	e.GET("/provinces", handler.GetAllProvince)
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
