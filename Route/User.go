package Route

import (
	"github.com/labstack/echo/v4"
	"laundry/Controller"
	"laundry/Middleware"
)

func UserRoute(e *echo.Echo) {
	api := e.Group("/user")

	// for authentication route
	api.POST("/login", Controller.Login)

	// protected route
	api.Use(Middleware.UserMiddleware)
}
