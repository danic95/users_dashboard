package v1

import (
	"github.com/labstack/echo/v4"
)

// RegisterRoutes initializes api v1 routes
func (h *Handler) RegisterRoutes(e *echo.Group) {

	// Swagger routes
	e.GET("/v1/docs", h.getSwaggerIndex)
	e.GET("/v1/docs/swagger.yml", h.getSwaggerSchema)

	// swagger:route GET /user users getUserRQ
	//
	// Retrieves an user
	//
	// Retrieves an user
	//
	// responses:
	//    200: getUserRS
	//    400: badRequestRS
	//	  500: serverErrorRS
	e.GET("/v1/user/:id", h.getUser)

	
	// swagger:route GET /friends friends getfriendsRQ
	//
	// Retrieves all friends
	//
	// Retrieves all fruebds
	//
	// responses:
	//    200: getFriendsRS
	//    400: badRequestRS
	//	  500: serverErrorRS
	e.GET("/v1/friends/:id", h.getFriends)
}
