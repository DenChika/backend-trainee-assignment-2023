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
		segment.POST("", h.CreateSegment)
		segment.DELETE("", h.DeleteSegment)
	}
	userSegment := router.Group("/user-segment")
	{
		userSegment.POST("", h.AddUserToSegment)
		userSegment.GET("", h.GetUserSegments)
	}
	return router
}
