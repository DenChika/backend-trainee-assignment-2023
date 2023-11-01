package handlers

import (
	"backend-trainee-assignment-2023/pkg/services"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	segment := router.Group("/segment", h.userIdentity)
	{
		segment.POST("/", h.CreateSegment)
		segment.DELETE("/", h.DeleteSegment)
	}

	userSegment := router.Group("/users-segments", h.userIdentity)
	{
		userSegment.POST("/", h.ManageUserToSegments)
		userSegment.GET("/", h.GetUserSegments)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signIn)
	}
	return router
}
