package routes

import (
	"github.com/labstack/echo/v4"
	"sejutaCita/services/user-service/delivery/handler/user"
	"sejutaCita/services/user-service/delivery/middleware"
)

func UserPath(e *echo.Echo, uh *user.UserHandler) {
	e.POST("create/user", uh.CreateUser(), middleware.JWTMiddleware())
	e.POST("/user/login", uh.LoginUserHandler())
	e.GET("/user", uh.GetUserById(), middleware.JWTMiddleware())
}
