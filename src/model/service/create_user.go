package service

import (
	"context"
	"fmt"
	"sarracena_erp/src/configuration/database"
	logger "sarracena_erp/src/configuration/logs"
	"sarracena_erp/src/configuration/rest_err"
	"sarracena_erp/src/model"
	"sarracena_erp/src/model/repository"

	"go.uber.org/zap"
)

// CreateUser processes the user creation logic, including password encryption and database interaction.
func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	// Encrypt the password
	userDomain.EncryptPassword()

	// Log the encrypted password (for debug purposes, avoid this in production logs)
	fmt.Println(userDomain.GetPassword())

	// Save the user to the database
	repo := repository.NewUserRepository(database.GetDB()) // Obtenha o pool de conexão
	createdUser, err := repo.CreateUser(context.Background(), userDomain)
	if err != nil {
		logger.Error("Error saving user to database", err)


		return rest_err.NewInternalServerError("Erro ao salvar o usuário no banco de dados")
	}

	logger.Info("User created successfully", zap.String("userId", createdUser.GetID()))

	return nil
}
