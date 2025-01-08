// Package controller provides interfaces and implementations for handling user-related requests.
package controller

import (
	"sarracena_erp/src/model/service"

	"github.com/gin-gonic/gin"
)

// NewUserControllerInterface creates a new instance of UserControllerInterface.
// Parameters:
// - serviceInterface: The service layer interface for user domain operations.
// Returns: A UserControllerInterface instance.
func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

// UserControllerInterface defines the methods for handling user-related HTTP requests.
type UserControllerInterface interface {
	// FindUserByID fetches a user by their ID.
	FindUserByID(c *gin.Context)
	
	// FindUserByEmail fetches a user by their email.
	FindUserByEmail(c *gin.Context)
	
	// DeleteUser deletes a user by their ID.
	DeleteUser(c *gin.Context)
	
	// CreateUser creates a new user.
	CreateUser(c *gin.Context)
	
	// UpdateUser updates an existing user's details by their ID.
	UpdateUser(c *gin.Context)
}

// userControllerInterface is the concrete implementation of UserControllerInterface.
type userControllerInterface struct {
	// service is the service layer interface for user domain operations.
	service service.UserDomainService
}
