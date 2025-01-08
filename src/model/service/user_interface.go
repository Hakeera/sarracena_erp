// Package service defines the interfaces and implementations for user operations.
package service

import (
	"sarracena_erp/src/configuration/rest_err"
	"sarracena_erp/src/model"
)

// NewUserDomainService creates a new instance of UserDomainService.
// Returns: An instance of UserDomainService.
func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

// UserDomainService defines the methods for user-related business logic.
type UserDomainService interface {
	// CreateUser creates a new user in the system.
	CreateUser(model.UserDomainInterface) *rest_err.RestErr

	// UpdateUser updates user information by their ID.
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr

	// FindUser retrieves a user by their ID.
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)

	// DeleteUser removes a user from the system by their ID.
	DeleteUser(string) *rest_err.RestErr
}

// userDomainService is the concrete implementation of UserDomainService.
type userDomainService struct{}
