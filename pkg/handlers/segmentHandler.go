package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Segment struct {
	Slug string
}

func CreateSegment(ctx echo.Context) error {
	var segment Segment
	err := ctx.Bind(&segment)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}
	return ctx.JSON(http.StatusOK, nil)
}
