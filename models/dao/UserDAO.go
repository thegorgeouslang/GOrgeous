// Author: James Mallon <jamesmallondev@gmail.com>
// dao package -
package dao

import (
	. "GoAuthorization/libs/databases"
	//log "GoAuthorization/libs/logger"
	. "GoAuthorization/models"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

// Struct type userDAO -
type userDAO struct{}

// UserDAO function -
func UserDAO() *userDAO {
	return &userDAO{}
}

// Create method - Stores a new user in the system
func (this *userDAO) Create(user *User) (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// it creates a hashed byte slice from the user password
		pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = pass // apply the encrypted pass to the user obj
		DbConn.GetConn().Create(&user)
		e = DbConn.Check() // check for errors
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
		DbConn.GetConn().Take(&user)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// GetUsers method -
func (this *userDAO) GetUsers() (users []User, e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		DbConn.GetConn().Find(&users)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}
