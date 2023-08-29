package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Segment struct {
	Slug string `param:"slug"`
}

func (h *Handler) CreateSegment(ctx echo.Context) error {
	var segment Segment
	if err := ctx.Bind(&segment); err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}
	id, err := h.service.Segment.Create(segment.Slug)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
