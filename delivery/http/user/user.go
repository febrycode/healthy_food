package httpuser

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fsetiawan29/healthy_food/domain/user"
	middlewareCustom "github.com/fsetiawan29/healthy_food/middleware"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	e.GET("health_check", handler.HealthCheck)
	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)

	// Restricted group
	r := e.Group("/profile")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &middlewareCustom.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("", handler.Profile)
	r.PUT("", handler.UpdateUser)
}

// Login will find the user authentication
func (u *UserHandler) Login(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var userParam models.User
	if err := c.Bind(&userParam); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	user, err := u.userUsecase.GetUserByEmail(ctx, userParam.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	if user.ID <= 0 {
		return c.JSON(http.StatusUnauthorized, models.ResponseJSON(http.StatusUnauthorized, "Email is not valid"))
	}

	if !middlewareCustom.ComparePassword(user.Password, middlewareCustom.GetPassword(userParam.Password)) {
		return c.JSON(http.StatusUnauthorized, models.ResponseJSON(http.StatusUnauthorized, "Email and password is incorrect"))
	}

	// Set custom claims
	claims := &middlewareCustom.JwtCustomClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	return c.JSON(http.StatusOK, models.ResponseToken(http.StatusOK, t, user.IsAdmin))
}

func (u *UserHandler) Register(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userParam := &models.User{}
	if err = c.Bind(userParam); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	user, err := u.userUsecase.GetUserByEmail(ctx, userParam.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	// Check email has been created or not
	// Check unique email
	if user.ID > 0 {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Email has been created"))
	}

	userParam.Password = middlewareCustom.HashAndSalt(middlewareCustom.GetPassword(userParam.Password))
	err = u.userUsecase.CreateUser(ctx, userParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	return c.JSON(http.StatusCreated, models.ResponseJSON(http.StatusCreated, "User created successfully"))
}

func (u *UserHandler) Profile(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userParam := c.Get("user").(*jwt.Token)
	claims := userParam.Claims.(*middlewareCustom.JwtCustomClaims)
	userID := claims.UserID

	user, err := u.userUsecase.GetUserByUserID(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) UpdateUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userParam := &models.User{}
	if err = c.Bind(userParam); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	userParamJWT := c.Get("user").(*jwt.Token)
	claims := userParamJWT.Claims.(*middlewareCustom.JwtCustomClaims)
	userID := claims.UserID

	if userParam.ID != userID {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	userData, err := u.userUsecase.GetUserByUserID(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	if userData.ID <= 0 {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	err = u.userUsecase.UpdateUser(ctx, userParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseJSON(http.StatusBadRequest, "Bad Request"))
	}

	return c.JSON(http.StatusCreated, models.ResponseJSON(http.StatusCreated, "User updated successfully"))
}

func (u *UserHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "success")
}
