// Author: James Mallon <jamesmallondev@gmail.com>
// dao package -
package dao

import (
	. "GoAuthentication/models"
	"errors"
)

// Struct type userDAO -
type userDAO struct{}

var UserDAO *userDAO

// variable simulates a database table
var UserTbl []User

// init function - data and process initialization
func init() {
	UserDAO = &userDAO{}
}

// Insert method -
func (this *userDAO) Insert(user *User) (e error) {
	if this.CheckExists(user.Email) {
		e = errors.New("User already exists in our database")
		return
	}
	user.Id = (len(UserTbl) + 1)     // get array length
	UserTbl = append(UserTbl, *user) // append a new user to the array (fake db table)
	return
}

// CheckExistent method -
func (this *userDAO) CheckExists(email string) bool {
	for _, user := range UserTbl {
		if user.Email == email {
			return true
		}
	}
	return false
}

// GetUser method -
func (this *userDAO) GetUser(id int) User {
	return UserTbl[id]
}

// GetUsers method -
func (this *userDAO) GetUsers() []User {
	return UserTbl
}

// GetPassByEmail method -
func (this *userDAO) GetByEmail(email string) User {
	for _, user := range UserTbl {
		if user.Email == email {
			return user
		}
	}
	return User{}
}
