// Sessionor: James Mallon <jamesmallondev@gmail.com>
// dao package - contain the system's daos
package dao

import (
	. "GoAuthorization/models"
	"time"
)

// Struct type sessionDAO -
type sessionDAO struct{}

var SessionDAO *sessionDAO
var SessionDB map[string]Session

// init function - data and process initialization
func init() {
	SessionDAO = &sessionDAO{}
	SessionDB = map[string]Session{}
}

// Insert method -
func (this *sessionDAO) Insert(sid string, session *Session) {
	SessionDB[sid] = *session
}

// GetSession method -
func (this *sessionDAO) GetSession(sid string) Session {
	return SessionDB[sid]
}

// GetSessions method -
func (this *sessionDAO) GetSessions() map[string]Session {
	return SessionDB
}

// Delete method -
func (this *sessionDAO) Delete(sid string) {
	delete(SessionDB, sid)
}

// Renew method -
func (this *sessionDAO) Renew(sid string) {
	session := SessionDB[sid]
	session.LastActivity = time.Now()
	SessionDB[sid] = session
}
