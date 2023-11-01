package handlers

import (
	"backend-trainee-assignment-2023/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary signUp
// @Tags auth
// @Description sign up
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body models.AuthRequest true "username, password"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(ctx echo.Context) error {
	var req models.AuthRequest
	if err := ctx.Bind(&req); err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	if err := h.service.Authorization.SignUp(req.Username, req.Password); err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
	})
}

// @Summary signIn
// @Tags auth
// @Description sign in
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body models.AuthRequest true "username, password"
// @Success 200 {object} models.SignInResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [get]
func (h *Handler) signIn(ctx echo.Context) error {
	var req models.AuthRequest
	if err := ctx.Bind(&req); err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	token, err := h.service.Authorization.SignIn(req.Username, req.Password)
	if err != nil {
		return newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, models.SignInResponse{Token: token})
}
