package service

import (
	"fmt"
	logger "sarracena_erp/src/configuration/logs"
	"sarracena_erp/src/configuration/rest_err"
	"sarracena_erp/src/model"

	"go.uber.org/zap"
)


func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}