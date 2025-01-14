// Package controller contains handlers for processing incoming HTTP requests.
package controller

import (
	"net/http"
	logger "sarracena_erp/src/configuration/logs"
	"sarracena_erp/src/configuration/validation"
	"sarracena_erp/src/model"
	"sarracena_erp/src/model/request"
	"sarracena_erp/src/view"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateUser handles the creation of a new user.
// It binds the request payload to a UserRequest struct, validates it, and processes the creation.
// Parameters:
// - c: The Gin context containing the HTTP request and response.
func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	// Bind JSON payload to the UserRequest struct and validate.
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	// Create a UserDomain instance from the request data.
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	// Não é necessário inicializar o repositório ou o serviço novamente
	// O serviço já foi instanciado em main.go e passado para o controlador

	// Criar o usuário
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	// Respond with the created user's data.
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))
}
