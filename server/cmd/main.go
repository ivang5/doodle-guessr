package main

import (
	"github.com/ivang5/doodle-guessr/server/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	addr := "localhost:3000"

	e := echo.New()
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.DEBUG)
	e.HideBanner = true

	r := router.Default(e)

	log.Fatal(r.Start(addr))
}
