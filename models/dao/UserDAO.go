// Author: James Mallon <jamesmallondev@gmail.com>
// dao package -
package dao

import (
	. "GoAuthentication/models"
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
func (this *userDAO) Insert(user *User) bool {
	user.Id = len(UserTbl)           // get array length
	UserTbl = append(UserTbl, *user) // append a new user to the array (fake db table)
	return true
}

// CheckExistent method -
func (this *userDAO) CheckExists(user *User, email string) bool {
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
