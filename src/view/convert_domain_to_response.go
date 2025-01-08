package view

import (
	"sarracena_erp/src/model"
	"sarracena_erp/src/model/response"
)

// ConvertDomainToResponse transforma um domínio de usuário em uma estrutura JSON 
// para ser enviada ao cliente.
// Recebe um objeto que implementa a interface UserDomainInterface e retorna um UserResponse.
func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
