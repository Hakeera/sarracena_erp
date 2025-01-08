// Package response defines the structures for outgoing HTTP responses.
package response

// UserResponse represents the structure of a user response object.
// Contains user details to be sent to the client.
type UserResponse struct {
	Email    string `json:"email"`    // User's email address.
	Password string `json:"password"` // User's hashed password.
	Name     string `json:"name"`     // User's name.
	Age      int8   `json:"age"`      // User's age.
}
