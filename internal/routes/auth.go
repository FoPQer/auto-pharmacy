package routes

import (
	"auto-pharmacy/internal/controllers"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func RegisterAuthRoutes(e *echo.Group, lc *controllers.LoginController) {
	e.POST("/login", lc.LoginHandler, middleware.BasicAuth(controllers.BasicAuth))
	e.POST("/register", lc.RegisterHandler)
}
