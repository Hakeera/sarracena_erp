package view

import (
	"sarracena_erp/src/model"
	"sarracena_erp/src/model/response"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
 ) response.UserResponse {
	return response.UserResponse{
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),	
	}
}