package main

import (
	"fmt"
	"os"

	"github.com/ggalihpp/shopit/bank"
	"github.com/ggalihpp/shopit/courier"
	"github.com/ggalihpp/shopit/item"
	mw "github.com/ggalihpp/shopit/middleware"
	"github.com/ggalihpp/shopit/user"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func setupHandler(e *echo.Echo) {
	connectionString := os.Getenv("GORM_CONNECTION")
	db, err := gorm.Open("postgres", connectionString)
	if err != nil || db.Error != nil {
		panic(err)
	}
	if os.Getenv("GORM_LOG") == "true" {
		db.LogMode(true)
	}

	//initDB(db)

	// bankCon := bank.Connection(db)
	// courierCon := courier.Connection(db)

	itemCon := item.Connection(db)
	userCon := user.Connection(db)

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	ig := e.Group("/item")
	itemHandlers := &item.Handler{
		Repository: &itemCon,
	}

	ig.Use(mw.CheckJWT)
	itemHandlers.SetRoutes(ig)

	ug := e.Group("/u")
	userHandlers := &user.Handler{
		Repository: &userCon,
	}

	userHandlers.SetRoutes(ug)
}

func initDB(db *gorm.DB) {
	if !db.HasTable(&item.Item{}) {
		db.Exec("CREATE EXTENSION zombodb;")
	}

	db.AutoMigrate(
		&user.Address{},
		&user.BankAccount{},
		&user.Order{},
		&user.OrderItem{},
		&bank.Bank{},
		&user.User{},
		&courier.Courier{},
		&item.Item{},
	)

	db.Exec(fmt.Sprintf("CREATE INDEX idxitems ON items USING zombodb ((items.*)) WITH (url='%s');", os.Getenv("ES_URL")))
}
