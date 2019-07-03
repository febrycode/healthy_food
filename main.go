package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

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

	userRepository := _userRepo.NewMysqlUserRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userUsecase := _userUsecase.NewUserUsecase(userRepository, timeoutContext)
	_userHttpDeliver.NewUserHandler(e, userUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
