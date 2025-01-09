// Package model defines the structures and methods for the user domain.
package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

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

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

// userDomain is the concrete implementation of UserDomainInterface.
type userDomain struct {
	ID		 string
	Email    string `json: "email"`
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

// GetEmail retrieves the user's email.
func (ud *userDomain) GetEmail() string {
	return ud.Email
}

// GetPassword retrieves the user's password.
func (ud *userDomain) GetPassword() string {
	return ud.Password
}

// GetName retrieves the user's name.
func (ud *userDomain) GetName() string {
	return ud.Name
}

// GetAge retrieves the user's age.
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

// EncryptPassword hashes the user's password using MD5.
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
