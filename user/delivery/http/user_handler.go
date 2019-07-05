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
	userUsecase user.Usecase
}

// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(e *echo.Echo, userUsecase user.Usecase) {
	handler := &UserHandler{
		userUsecase: userUsecase,
	}

	e.POST("/register", handler.Register)
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

	user, err := u.userUsecase.GetUserByEmail(ctx, userParam.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) Register(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userParam := &models.User{}
	if err = c.Bind(userParam); err != nil {
		return err
	}

	err = u.userUsecase.CreateUser(ctx, userParam)
	if err != nil {
		return err
	}

	return nil
}
