package http

import (
	"context"
	"net/http"

	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/user"
	"github.com/labstack/echo"
)

// UserHandler represent the httphandler for user
type UserHandler struct {
	UserUsecase user.Usecase
}

// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(e *echo.Echo, userUsecase user.Usecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}

	e.POST("/login", handler.Login)
}

// Login will find the user authentication
func (u *UserHandler) Login(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var userParam models.User
	err := c.Bind(&userParam)
	if err != nil {
		return err
	}

	user, err := u.UserUsecase.GetUserByEmail(ctx, userParam.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)

}
