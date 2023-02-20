package main

import (
	"github.com/labstack/echo"
	"github.com/rvldodo/go-echo-tut/cmd/api/handlers"
)

func main() {
	e := echo.New()

	// routes
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":3000"))

}