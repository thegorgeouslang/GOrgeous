// Author: James Mallon <jamesmallondev@gmail.com>
// dao package -
package dao

import (
	. "GoAuthorization/models"
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
	UserTbl = append(UserTbl, User{})
}

// Insert method -
func (this *userDAO) Create(user *User) (e error) {
	udb := this.GetByEmail(user.Email)
	if len(udb.Email) > 0 {
		e = errors.New("User already exists in our database")
		return
	}
	user.Id = len(UserTbl)
	UserTbl = append(UserTbl, *user)
	return
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
