package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rhidayat1980/fitness-api/handlers"
	"github.com/rhidayat1980/fitness-api/storage"
)

func main() {

	e := echo.New()

	e.GET("/", handlers.Home)
	storage.InitDB()

	// e.Use(handlers.LogRequest)
	e.Use(middleware.Logger(), middleware.Recover())

	e.POST("/users", handlers.CreateUser)
	e.PUT("/user/:id", handlers.UpdateUser)
	e.POST("/measurements", handlers.CreateMeasurement)
	e.PUT("/measurement/:id", handlers.UpdateMeasurement)

	e.Logger.Fatal(e.Start(":8080"))
}
