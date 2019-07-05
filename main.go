package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_foodHttpDeliver "github.com/febrycode/healthy_food/food/delivery/http"
	_foodRepo "github.com/febrycode/healthy_food/food/repository"
	_foodUsecase "github.com/febrycode/healthy_food/food/usecase"
	_foodDetailRepo "github.com/febrycode/healthy_food/food_detail/repository"
	"github.com/febrycode/healthy_food/middleware"
	_userHttpDeliver "github.com/febrycode/healthy_food/user/delivery/http"
	_userRepo "github.com/febrycode/healthy_food/user/repository"
	_userUsecase "github.com/febrycode/healthy_food/user/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
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
	userUsecase := _userUsecase.NewUserUsecase(userRepository, timeoutContext)
	_userHttpDeliver.NewUserHandler(e, userUsecase)

	foodRepository := _foodRepo.NewMysqlFoodRepository(dbConn)
	foodDetailRepository := _foodDetailRepo.NewMysqlFoodDetailRepository(dbConn)
	foodUsecase := _foodUsecase.NewFoodUsecase(foodRepository, foodDetailRepository, timeoutContext)
	_foodHttpDeliver.NewFoodHandler(e, foodUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
