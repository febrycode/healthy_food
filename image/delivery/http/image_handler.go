package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fsetiawan29/healthy_food/image"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/labstack/echo"
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

	e.POST("/image", handler.CreateImage)

	// Restricted group
	// r := e.Group("/image")

	// // Configure middleware with the custom claims type
	// config := middleware.JWTConfig{
	// 	Claims:     &middlewareCustom.JwtCustomClaims{},
	// 	SigningKey: []byte("secret"),
	// }

	// r.Use(middleware.JWTWithConfig(config))
	// r.POST("", handler.CreateImage)
}

func (i *ImageHandler) CreateImage(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	images := form.File["images"]

	var result []string

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
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		r := c.Request()
		fileName := c.Scheme() + "://" + r.Host + "/upload/" + image.Filename
		err = i.imageUsecase.CreateImage(ctx, &models.Image{
			Name: fileName,
		})
		if err != nil {
			return err
		}
		result = append(result, fileName)

	}

	return c.JSON(http.StatusCreated, models.ResponseImage(http.StatusCreated, result))
}
