package handlers

import (
	"backend-trainee-assignment-2023/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AddUserToSegmentRequest struct {
	SlugsToAdd    []string `json:"slugs-to-add" validate:"empty=false"`
	SlugsToRemove []string `json:"slugs-to-remove" validate:"empty=false"`
	UserId        uint     `json:"user-id" validate:"empty=false"`
}

type GetUserSegmentsRequest struct {
	UserId uint `json:"user-id" validate:"empty=false"`
}

func (h *Handler) AddUserToSegment(ctx echo.Context) error {
	var req AddUserToSegmentRequest
	if err := ctx.Bind(&req); err != nil {
		return helpers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	if err := h.service.AddUserToSegment(req.SlugsToAdd, req.SlugsToRemove, req.UserId); err != nil {
		return helpers.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
	})
}

func (h *Handler) GetUserSegments(ctx echo.Context) error {
	var req GetUserSegmentsRequest
	if err := ctx.Bind(&req); err != nil {
		return helpers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	segments, err := h.service.User.GetUserSegments(req.UserId)
	if err != nil {
		return helpers.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"segments": segments,
		"status":   http.StatusOK,
	})
}
