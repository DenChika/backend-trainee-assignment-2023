package services

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
		return ctx.String(http.StatusBadRequest, "Bad request")
	}
	return ctx.String(http.StatusOK, "")
}
