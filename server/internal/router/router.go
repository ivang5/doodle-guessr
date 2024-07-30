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

	apiRoute.GET("/highscore", handlers.ReadHighscores)
	apiRoute.POST("/highscore", handlers.SetHighscore)

	return r
}

func (r *Router) Start(addr string) error {
	return r.echo.Start(addr)
}
