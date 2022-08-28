package routes

import (
	"github.com/labstack/echo/v4"
	"sejutaCita/services/admin-service/delivery/handler/admin"
	"sejutaCita/services/admin-service/delivery/middleware"
)

func AdminPath(e *echo.Echo, ah *admin.AdminHandler) {
	e.POST("/create/admin", ah.CreateAdmin())
	e.POST("/login", ah.LoginHandler())
	e.GET("/admin", ah.GetAdminById(), middleware.JWTMiddleware())
	e.PUT("/update", ah.UpdateAdmin(), middleware.JWTMiddleware())
	e.DELETE("/delete", ah.DeleteAdmin(), middleware.JWTMiddleware())
}
