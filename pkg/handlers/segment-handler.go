package handlers

import (
	"backend-trainee-assignment-2023/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SegmentRequest struct {
	Slug string
}

func (h *Handler) CreateSegment(ctx echo.Context) error {
	var req SegmentRequest
	if err := ctx.Bind(&req); err != nil {
		return helpers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	id, err := h.service.Segment.Create(req.Slug)
	if err != nil {
		return helpers.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":     id,
		"status": http.StatusOK,
	})
}

func (h *Handler) DeleteSegment(ctx echo.Context) error {
	var req SegmentRequest
	if err := ctx.Bind(&req); err != nil {
		return helpers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	if err := h.service.Segment.Delete(req.Slug); err != nil {
		return helpers.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
	})
}
