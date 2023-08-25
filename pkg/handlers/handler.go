package handlers

import "github.com/labstack/echo/v4"

func InitRoutes() *echo.Echo {
	router := echo.New()
	segment := router.Group("/segment")
	{
		segment.POST("/", func(ctx echo.Context) error { return nil })
		segment.GET("/", func(ctx echo.Context) error { return nil })
	}
	return router
}
