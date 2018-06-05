package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/y-ogura/yakiniku/db"
	"github.com/y-ogura/yakiniku/db/migration"
)

var database *gorm.DB

func setup(e *echo.Echo) {
	// connect database
	database = db.ConnectDB()
	// run migrate
	migration.Execute()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
