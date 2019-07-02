// Author: James Mallon <jamesmallondev@gmail.com>
// models package - contains the system's models
package models

import (
	//	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	//	gorm.Model
	Email        string
	LastActivity time.Time
}
