package main

import (
	"log"
	"sarracena_erp/src/configuration/database"
	logger "sarracena_erp/src/configuration/logs"
	"sarracena_erp/src/controller"
	"sarracena_erp/src/controller/routes"
	"sarracena_erp/src/model/repository"
	"sarracena_erp/src/model/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	// Carregar Variáveis de Ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conexão com Banco de Dados
	database.InitDatabase()
	defer database.CloseDatabase()

	// Inicializar Repositório
	db := database.GetDB()
	userRepo := repository.NewUserRepository(db)

	// Inicializar Serviço
	userService := service.NewUserDomainService(userRepo)

	// Inicializar Controlador
	userController := controller.NewUserControllerInterface(userService)

	// Configurar Roteador
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	// Iniciar Servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
