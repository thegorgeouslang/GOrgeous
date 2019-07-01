// Author: James Mallon <jamesmallondev@gmail.com>
// models package - contains the models of the app
package models

type User struct {
	Id       int
	Email    string
	Password []byte
	Role     string
}
