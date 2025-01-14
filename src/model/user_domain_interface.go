package model


// UserDomainInterface defines the methods for the user domain object.
type UserDomainInterface interface {
	// GetEmail retrieves the user's email.
	GetEmail() string
	// GetPassword retrieves the user's password.
	GetPassword() string
	// GetAge retrieves the user's age.
	GetAge() int8
	// GetName retrieves the user's name.
	GetName() string
	// EncryptPassword hashes the user's password.

	SetID(string)
	GetJSONValue() (string, error)

	EncryptPassword()

	GetID() string

}

// NewUserDomain creates a new user domain instance.
// Parameters:
// - email: The user's email address.
// - password: The user's password.
// - name: The user's name.
// - age: The user's age.
// Returns: A UserDomainInterface instance.
func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Email: email,
		Password: password,
		Name: name,
		Age: age,
	}
}