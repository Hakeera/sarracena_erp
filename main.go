package main

import (
	"log"
	logger "sarracena_erp/src/configuration/logs"
	"sarracena_erp/src/controller"
	"sarracena_erp/src/controller/routes"
	"sarracena_erp/src/model/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func main() {
    logger.Info("About to start user application")
    
    err:= godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
	//Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

    router := gin.Default()

    routes.InitRoutes(&router.RouterGroup, userController)

    if err:= router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}


