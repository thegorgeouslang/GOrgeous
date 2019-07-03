// Sessionor: James Mallon <jamesmallondev@gmail.com>
// dao package - contain the system's daos
package dao

import (
	. "GoAuthorization/libs/databases"
	//log "GoAuthorization/libs/logger"
	. "GoAuthorization/models"
	"sync"
	"time"
)

// Struct type sessionDAO -
type sessionDAO struct{}

// SessionDAO function -
func SessionDAO() *sessionDAO {
	return &sessionDAO{}
}

// Insert method -
func (this *sessionDAO) Create(s *Session) (e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		DbConn.GetConn().Create(&s)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// GetSession method -
func (this *sessionDAO) GetSession(sid string) (s Session, e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		s = Session{SID: sid}
		DbConn.GetConn().Last(&s) // searching from the last (natural order of insertion)
		e = DbConn.Check()        // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// GetSessions method -
func (this *sessionDAO) GetSessions() (sessions []Session, e error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		DbConn.GetConn().Find(&sessions)
		e = DbConn.Check() // check for errors
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// Delete method -
func (this *sessionDAO) Remove(sid string) (e error) {
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
func (this *sessionDAO) Renew(sid string) (e error) {
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
