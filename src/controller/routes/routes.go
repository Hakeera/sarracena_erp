// Package routes defines the API routes and their associated handlers.
package routes

import (
	"sarracena_erp/src/controller"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes the API routes for user-related operations.
// Parameters:
// - r: The Gin RouterGroup to register the routes.
// - userController: The controller interface implementing user operations.
func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", userController.FindUserByID)       // Fetches a user by their ID.
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail) // Fetches a user by their email.
	r.POST("/createUser", userController.CreateUser)                // Creates a new user.
	r.PUT("/updateUser/:userId", userController.UpdateUser)         // Updates an existing user by their ID.
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)      // Deletes a user by their ID.
}
