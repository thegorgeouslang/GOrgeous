// Author: James Mallon <jamesmallondev@gmail.com>
// model package -
package models

import (
	. "TheGorgeous/libs/databases"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

// Struct type User -
type User struct {
	gorm.Model
	Email    string
	Password []byte
	Role     string
}

// NewUser function -
func NewUser() *User {
	return &User{}
}

// Create method - Stores a new user in the system
func (this *User) Create() (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		_ = this.GetByEmail()
		if this.ID > 0 {
			e = errors.New("The user already exists in the database")
		} else {
			// it creates a hashed byte slice from the user password
			pass, _ := bcrypt.GenerateFromPassword([]byte(this.Password), bcrypt.DefaultCost)
			this.Password = pass // apply the encrypted pass to the user obj
			DbConn.GetConn().Create(&this)
			e = DbConn.Check() // check for errors
		}
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// GetByEmail method -
func (this *User) GetByEmail() (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		DbConn.GetConn().Where("email = ?", this.Email).Find(&this)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// GetUsers method -
func (this *User) GetUsers() (users []User, e error) {
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

// DeleteUser method -
func (this *User) DeleteUser(u User) (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		DbConn.GetConn().Delete(&u)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}
