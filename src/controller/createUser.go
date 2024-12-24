package controller

import (
	"net/http"

	logger "sarracena_erp/src/configuration/logs"
	"sarracena_erp/src/configuration/validation"
	"sarracena_erp/src/model/request"
	"sarracena_erp/src/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	// Vincula o JSON ao struct e verifica se há erros
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		Email: userRequest.Email,
		Password:    "test",
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
