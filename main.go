package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"github.com/suryakun/skeleton-go/middleware"
	delivery "github.com/suryakun/skeleton-go/user/delivery/http"
	_userRepository "github.com/suryakun/skeleton-go/user/repository"
	"github.com/suryakun/skeleton-go/user/service"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service run on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	params := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
	db, err := gorm.Open("postgres", params)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()
	middle := middleware.InitMiddleware()
	e.Use(middle.CORS)
	_userRepo := _userRepository.NewUserRepository(db)

	userService := service.NewUserService(
		_userRepo,
	)
	delivery.NewUserHandler(e, userService)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
