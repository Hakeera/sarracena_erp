// Package service contains the business logic for user-related operations.
package service

import (
	"fmt"
	logger "sarracena_erp/src/configuration/logs"
	"sarracena_erp/src/configuration/rest_err"
	"sarracena_erp/src/model"

	"go.uber.org/zap"
)

// CreateUser processes the user creation logic, including password encryption.
// Parameters:
// - userDomain: The user domain object containing user details.
// Returns: A RestErr instance in case of errors, or nil on success.
func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
