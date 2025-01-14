package repository

import (
	"context"
	"sarracena_erp/src/configuration/rest_err"
	"sarracena_erp/src/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	databaseConnection *pgxpool.Pool
}

// Definição da interface UserRepository
type UserRepository interface {
	// Recebe context.Context como primeiro parâmetro
	CreateUser(
		ctx context.Context,
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}

// NewUserRepository cria uma nova instância do repositório de usuário.
// Recebe: *pgxpool.Pool como conexão com o banco de dados.
// Retorna: UserRepository.
func NewUserRepository(database *pgxpool.Pool) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}
