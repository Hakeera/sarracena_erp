package service

import (
	"sarracena_erp/src/configuration/rest_err"
	"sarracena_erp/src/model"
	"sarracena_erp/src/model/repository"
)

// NewUserDomainService cria uma nova instância de UserDomainService.
// Agora recebe o UserRepository como parâmetro.
func NewUserDomainService(userRepo repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepo: userRepo, // Armazena o repositório no serviço
	}
}

// UserDomainService define os métodos para a lógica de negócio relacionada a usuários.
type UserDomainService interface {
	// CreateUser cria um novo usuário no sistema.
	CreateUser(model.UserDomainInterface) *rest_err.RestErr

	// UpdateUser atualiza as informações de um usuário pelo ID.
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr

	// FindUser recupera um usuário pelo ID.
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)

	// DeleteUser remove um usuário do sistema pelo ID.
	DeleteUser(string) *rest_err.RestErr
}

// userDomainService é a implementação concreta da interface UserDomainService.
type userDomainService struct {
	userRepo repository.UserRepository
}

