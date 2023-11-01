package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		header, ok := ctx.Request().Header[authHeader]
		if !ok {
			return newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		}
		if len(header) > 1 {
			return newErrorResponse(ctx, http.StatusUnauthorized, "too much auth headers")
		}
		headerValue := strings.Split(header[0], " ")
		if len(headerValue) != 2 {
			return newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		}
		userId, err := h.service.Authorization.ParseToken(headerValue[1])
		if err != nil {
			return newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		}
		ctx.Set(userCtx, userId)
		return next(ctx)
	}
}

func getUserId(ctx echo.Context) (uint, error) {
	userId := ctx.Get(userCtx)
	if userId == 0 {
		return 0, errors.New("user id not found")
	}
	id, ok := userId.(uint)
	if !ok {
		return 0, errors.New("user id type is not valid")
	}
	return id, nil
}
