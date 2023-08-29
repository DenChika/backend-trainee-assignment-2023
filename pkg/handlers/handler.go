package handlers

import (
	"backend-trainee-assignment-2023/pkg/services"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()
	segment := router.Group("/segment")
	{
		segment.POST("/segment", h.CreateSegment)
		segment.GET("/", func(ctx echo.Context) error { return nil })
	}
	return router
}
