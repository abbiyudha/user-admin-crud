package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"sejutaCita/services/user-service/configs"
	handler "sejutaCita/services/user-service/delivery/handler/user"
	"sejutaCita/services/user-service/delivery/routes"
	repository "sejutaCita/services/user-service/repository/user"
	usecase "sejutaCita/services/user-service/usecase/user"
	"sejutaCita/services/user-service/utils"
)

func main() {
	config := configs.GetConfig()
	db, _ := utils.Connect(config)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	e := echo.New()
	routes.UserPath(e, userHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
