// Author: James Mallon <jamesmallondev@gmail.com>
// session package -
package session

import (
	conf "GoAuthentication/configs"
	. "GoAuthentication/models"
	. "GoAuthentication/models/dao"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"
)

var (
	sessCookie *http.Cookie
	sessExp    int
)

// Struct type session -
type sessionHelper struct{}

// init function - data and process initialization
func SessionHelper() *sessionHelper {
	sessExp, _ = strconv.Atoi(conf.Env["session_exp"])
	return &sessionHelper{}
}

// GetSession method -
func (this *sessionHelper) Start(w http.ResponseWriter, r *http.Request) string {
	var err error
	sessCookie, err = r.Cookie("session") // create the cookie
	if err != nil {
		sID := uuid.NewV4() // create the universal unique id
		sessCookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	sessCookie.MaxAge = sessExp
	http.SetCookie(w, sessCookie)
	return sessCookie.Value
}

//
func (this *sessionHelper) User(w http.ResponseWriter, r *http.Request) User {
	sessCookie.MaxAge = sessExp
	http.SetCookie(w, sessCookie)

	var user User
	// if the user exists already, get user
	session := SessionDAO.GetSession(sessCookie.Value) // retrieve the session
	if session.UserId > 0 {                            // check for the user id
		SessionDAO.Renew(sessCookie.Value)     // update LastActivity
		user = UserDAO.GetUser(session.UserId) // retrieve user
	}
	return user
}

// CloseSession method -
func (this *sessionHelper) Close(w http.ResponseWriter, r *http.Request) {
	sessCookie.MaxAge = -1
	http.SetCookie(w, sessCookie)
}

func (s *sessionHelper) CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	sessions := SessionDAO.GetSessions()
	for k, v := range sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			SessionDAO.Delete(k)
		}
	}
	fmt.Println("AFTER CLEAN") // for demonstration purposes
}
