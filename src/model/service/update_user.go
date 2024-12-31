package service

import (
	"sarracena_erp/src/configuration/rest_err"
	"sarracena_erp/src/model"
)


func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}