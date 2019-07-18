package http

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/fsetiawan29/healthy_food/image"
	middlewareCustom "github.com/fsetiawan29/healthy_food/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ImageHandler represent the httphandler for user
type ImageHandler struct {
	imageUsecase image.Usecase
}

// NewImageHandler will initialize the user/ resources endpoint
func NewImageHandler(e *echo.Echo, imageUsecase image.Usecase) {
	handler := &ImageHandler{
		imageUsecase: imageUsecase,
	}

	e.Static("/upload", "upload")

	// Restricted group
	r := e.Group("/image")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &middlewareCustom.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.POST("", handler.CreateImage)
}

func (i *ImageHandler) CreateImage(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	reference_id := c.FormValue("reference_id")

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	images := form.File["images"]

	fmt.Println(reference_id)

	fmt.Println(images)

	for _, image := range images {
		// Source
		src, err := image.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		if _, err := os.Stat("./upload"); os.IsNotExist(err) {
			err = os.Mkdir("./upload", os.ModePerm)
		}

		// Destination
		dst, err := os.Create(fmt.Sprintf("./upload/%s", image.Filename))
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return nil
}
