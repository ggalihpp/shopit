package main

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	envFile := os.Getenv("ENV")
	if envFile == "" {
		envFile = ".env"
	}

	// Load env data
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	}))

	setupHandler(e)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
