// Author: James Mallon <jamesmallondev@gmail.com>
// dao package -
package dao

import (
	. "GoAuthorization/libs/databases"
	log "GoAuthorization/libs/logger"
	. "GoAuthorization/models"
	"golang.org/x/crypto/bcrypt"
	"sync"
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

// Create method - Stores a new user in the system
func (this *userDAO) Create(user *User) (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// it creates a hashed byte slice from the user password
		pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = pass // apply the encrypted pass to the user obj

		if e = DbConn.Insert(&user); e != nil { // try to store the user in the db
			log.Write("error", e.Error(), log.Trace())
		}
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// GetUser method -
func (this *userDAO) GetByEmail(email string) (user User, e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		user = User{Email: email}
		if e = DbConn.SelectOne(&user); e != nil { // try to store the user in the db
			log.Write("error", e.Error(), log.Trace())
		}
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// GetUsers method -
func (this *userDAO) GetUsers() []User {
	return UserTbl
}
