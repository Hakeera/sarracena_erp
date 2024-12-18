package controller

import (
	"fmt"
	"sarracena_erp/src/configuration/rest_err"
	"sarracena_erp/src/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	// Vincula o JSON ao struct e verifica se há erros
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Campos incorretos: %v", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	// Exibe o conteúdo do usuário para fins de debug
	fmt.Println(userRequest)

	// Retorna sucesso com os dados do usuário recebido
	c.JSON(201, gin.H{
		"message": "Usuário criado com sucesso!",
		"data":    userRequest,
	})
}
