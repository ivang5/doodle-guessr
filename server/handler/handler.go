package handler

import "github.com/labstack/echo/v4"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Register(g *echo.Group) {
	g.POST("/predict", h.Predict)
	g.POST("/print", h.Print)
}
