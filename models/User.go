// Author: James Mallon <jamesmallondev@gmail.com>
// models package - contains the models of the app
package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password []byte
	Role     string
}
