package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/fsetiawan29/healthy_food/middleware"

	_foodRepo "github.com/fsetiawan29/healthy_food/repository/food"
	_foodDetailRepo "github.com/fsetiawan29/healthy_food/repository/food_detail"
	_imageRepo "github.com/fsetiawan29/healthy_food/repository/image"
	_provinceRepo "github.com/fsetiawan29/healthy_food/repository/province"
	_userRepo "github.com/fsetiawan29/healthy_food/repository/user"

	_foodUsecase "github.com/fsetiawan29/healthy_food/usecase/food"
	_imageUsecase "github.com/fsetiawan29/healthy_food/usecase/image"
	_provinceUsecase "github.com/fsetiawan29/healthy_food/usecase/province"
	_userUsecase "github.com/fsetiawan29/healthy_food/usecase/user"

	_foodHttpDeliver "github.com/fsetiawan29/healthy_food/delivery/http/food"
	_imageHttpDeliver "github.com/fsetiawan29/healthy_food/delivery/http/image"
	_provinceHttpDeliver "github.com/fsetiawan29/healthy_food/delivery/http/province"
	_userHttpDeliver "github.com/fsetiawan29/healthy_food/delivery/http/user"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.password`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sqlx.Open("mysql", dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middleware := middleware.InitMiddleware()
	e.Use(middleware.CORS)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepo.NewMysqlUserRepository(dbConn)
	foodRepository := _foodRepo.NewMysqlFoodRepository(dbConn)
	foodDetailRepository := _foodDetailRepo.NewMysqlFoodDetailRepository(dbConn)
	imageRepository := _imageRepo.NewMysqlImageRepository(dbConn)
	provinceRepository := _provinceRepo.NewMysqlProvinceRepository(dbConn)

	userUsecase := _userUsecase.NewUserUsecase(userRepository, timeoutContext)
	foodUsecase := _foodUsecase.NewFoodUsecase(foodRepository, foodDetailRepository, timeoutContext, imageRepository)
	provinceUsecase := _provinceUsecase.NewProvinceUsecase(provinceRepository, timeoutContext)
	imageUsecase := _imageUsecase.NewImageUsecase(imageRepository, timeoutContext)

	_userHttpDeliver.NewUserHandler(e, userUsecase)
	_foodHttpDeliver.NewFoodHandler(e, foodUsecase)
	_provinceHttpDeliver.NewProvinceHandler(e, provinceUsecase)
	_imageHttpDeliver.NewImageHandler(e, imageUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
