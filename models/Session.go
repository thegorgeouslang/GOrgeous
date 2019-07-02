// Author: James Mallon <jamesmallondev@gmail.com>
// models package - contains the system's models
package models

import (
	"time"
)

type Session struct {
	UserId       int
	Email        string
	LastActivity time.Time
}
