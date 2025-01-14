package model

import (
	"encoding/json"
	"fmt"
)

// userDomain is the concrete implementation of UserDomainInterface.
type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int8
}

// SetID sets the user's ID.
func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

// GetJSONValue returns the JSON representation of the user domain object.
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

func (ud *userDomain) GetID() string {
	return ud.ID
}