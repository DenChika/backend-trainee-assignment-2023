package handlers

import (
	"backend-trainee-assignment-2023/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary ManageUserToSegments
// @Security ApiKeyAuth
// @Tags users-segment
// @Description add and remove segments from user
// @ID manage-user-to-segments
// @Accept  json
// @Produce  json
// @Param input body models.ManageUserToSegmentsRequest true "slugs to add and remove, user id"
// @Success 200 {object} models.ManageUserToSegmentsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users-segments/ [post]
func (h *Handler) ManageUserToSegments(ctx echo.Context) error {
	userId, err := getUserId(ctx)
	if err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	var req models.ManageUserToSegmentsRequest
	if err := ctx.Bind(&req); err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	resp, err := h.service.ManageUserToSegments(req.SlugsToAdd, req.SlugsToRemove, userId)
	if err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, models.ManageUserToSegmentsResponse{
		SlugsHaveBeenAdded:   resp.SlugsHaveBeenAdded,
		SlugsHaveBeenRemoved: resp.SlugsHaveBeenRemoved,
	})
}

// @Summary GetUserSegments
// @Security ApiKeyAuth
// @Tags users-segment
// @Description get all user segments
// @ID get-user-segments
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetUserSegmentsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users-segments/ [get]
func (h *Handler) GetUserSegments(ctx echo.Context) error {
	userId, err := getUserId(ctx)
	if err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	slugs, err := h.service.UsersSegment.GetUserSegments(userId)
	if err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, models.GetUserSegmentsResponse{Slugs: slugs})
}
