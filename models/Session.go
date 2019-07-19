// Author: James Mallon <jamesmallondev@gmail.com>
// models package - contains the system's models
package models

import (
	. "TheGorgeous/libs/databases"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

// Struct type Session -
type Session struct {
	gorm.Model
	SID          string
	Email        string
	LastActivity time.Time
}

// NewSession function -
func NewSession() *Session {
	return &Session{}
}

// Insert method -
func (this *Session) Create() (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		DbConn.GetConn().Create(&this)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// GetSession method -
func (this *Session) GetSession(sid string) (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		this.SID = sid
		DbConn.GetConn().Last(&this) // searching from the last (natural order of insertion)
		e = DbConn.Check()           // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// GetSessions method -
func (this *Session) GetSessions() (Sessions []Session, e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		DbConn.GetConn().Find(&Sessions)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// Delete method -
func (this *Session) Remove(sid string) (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		s := Session{SID: sid}
		DbConn.GetConn().Delete(&s)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// Renew method -
func (this *Session) Renew(sid string) (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		s := Session{SID: sid}
		changes := map[string]interface{}{"LastActivity": time.Now()}
		DbConn.GetConn().Model(&s).Updates(changes)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}
