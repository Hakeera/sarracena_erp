package repository

import (
	"context"
	"fmt"
	"sarracena_erp/src/configuration/logs"
	"sarracena_erp/src/configuration/rest_err"
	"sarracena_erp/src/model"

	"go.uber.org/zap"
)

// Método para criar um usuário
func (ur *userRepository) CreateUser(
	ctx context.Context,
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	// Log de início da operação
	logger.Info("Init CreateUser repository")

	// Definição da consulta SQL para inserir o usuário no banco
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	// Capturando os dados do domínio
	name := userDomain.GetName()
	email := userDomain.GetEmail()
	password := userDomain.GetPassword()

	// Executando a consulta no banco de dados
	var id int
	err := ur.databaseConnection.QueryRow(ctx, query, name, email, password).Scan(&id)
	if err != nil {
		// Log de erro caso a inserção falhe
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	// Ajustando o ID do domínio para o valor retornado pelo banco
	userDomain.SetID(fmt.Sprintf("%d", id))

	// Retornando o domínio de usuário com o ID ajustado
	return userDomain, nil
}
