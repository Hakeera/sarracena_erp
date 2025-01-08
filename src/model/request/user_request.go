// Package request defines the structures for incoming HTTP requests.
package request

// UserRequest represents the structure of a user creation or update request.
// Fields include validation rules to ensure proper data input.
type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`               // User's email address.
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"` // User's password with validation.
	Name     string `json:"name" binding:"required,min=4,max=100"`        // User's name, with a length constraint.
	Age      int8   `json:"age" binding:"required,min=1,max=140"`         // User's age, within a valid range.
}
