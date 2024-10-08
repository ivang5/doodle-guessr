package router

import (
	"github.com/ivang5/doodle-guessr/server/internal/handlers"
	"github.com/labstack/echo/v4"
)

type Router struct {
	echo *echo.Echo
}

func New(e *echo.Echo) *Router {
	return &Router{
		echo: e,
	}
}

func Default(e *echo.Echo) *Router {
	r := New(e)
	r.echo.Static("", "./static")

	apiRoute := r.echo.Group("/api")

	apiRoute.POST("/predict", handlers.Predict)

	apiRoute.POST("/scores", handlers.SetScore)
	apiRoute.GET("/scores", handlers.ReadScores)

	return r
}

func (r *Router) Start(addr string) error {
	return r.echo.Start(addr)
}
