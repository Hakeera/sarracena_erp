package main

import (
	"log"
	logger "sarracena_erp/src/configuration/logs"
	"sarracena_erp/src/controller"
	"sarracena_erp/src/controller/routes"
	"sarracena_erp/src/model/database"
	"sarracena_erp/src/model/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func main() {
    logger.Info("About to start user application")
    
	// Carregar Variáveis de Ambiente
    err:= godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
	// Conexão com Banco de Dados
	database.InitDatabase()
	defer database.CloseDatabase()

	// Inicilizar Dependências
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	// Configurar Roteador
    router := gin.Default()
    routes.InitRoutes(&router.RouterGroup, userController)

	// Iniciar servidor
    if err:= router.Run(":8080"); err != nil {
        log.Fatal(err)
    }

	
}


