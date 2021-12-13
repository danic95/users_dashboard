package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sanservices/apicore/helper"
	"net/http"
)

func (h *Handler) getUser(c echo.Context) error {
	queryUserID := c.QueryParam("id")
	userID := queryUserID
	if userID == "" {
		return helper.RespondError(c, http.StatusBadRequest, errors.New("id query param must be string"))
	}

	user, err := h.service.GetUser(c.Request().Context(), userID)
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, errors.New("something went wrong"))
	}
	return helper.RespondOk(c, user)
}
