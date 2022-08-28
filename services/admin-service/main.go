package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"sejutaCita/services/admin-service/configs"
	handler "sejutaCita/services/admin-service/delivery/handler/admin"
	"sejutaCita/services/admin-service/delivery/routes"
	repository "sejutaCita/services/admin-service/repository/admin"
	usecase "sejutaCita/services/admin-service/usecase/admin"
	"sejutaCita/services/admin-service/utils"
)

func main() {
	config := configs.GetConfig()
	db, _ := utils.Connect(config)

	adminRepo := repository.NewAdminRepository(db)
	adminUsecase := usecase.NewAdminUsecase(adminRepo)
	adminHandler := handler.NewAdminHandler(adminUsecase)

	e := echo.New()
	routes.AdminPath(e, adminHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
