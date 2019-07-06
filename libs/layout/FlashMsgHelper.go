// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package layout

import (
	log "GoAuthorization/libs/logger"
	"encoding/base64"
	"net/http"
	"time"
)

// Struct type flashMessenger -
type FlashMsgHelper struct{}

// init function - data and process initialization
func init() {

}

//  Set function - it creates the cookie and set the flash message
func (this *FlashMsgHelper) Set(w *http.ResponseWriter, message string) {
	c := http.Cookie{ // creates the cookie
		Name:  "flashmessenger",
		Value: base64.URLEncoding.EncodeToString([]byte(message)),
	}
	http.SetCookie(*w, &c)
}

// Get method - it prints the message stored in the cookie
func (this *FlashMsgHelper) Get(w http.ResponseWriter, r *http.Request) (val string) {
	c, e := r.Cookie("flashmessenger")
	if e != nil { // check for errors - check if the cookie with the specific name is setted
		//log.Write("error", e.Error(), log.Trace())
		return
	} else {
		rc := http.Cookie{
			Name:    "flashmessenger",
			MaxAge:  -1,              // a negative number means it's already expired
			Expires: time.Unix(1, 0), // an expiration value that's in the past - already expired
		}
		cval, e := base64.URLEncoding.DecodeString(c.Value) // decode the value of the cookie
		val = string(cval)
		http.SetCookie(w, &rc) // override the values of the cookie setted in the setMessage function
		if e != nil {
			log.Write("Error", e.Error(), log.Trace())
			return
		}
		return
	}
}
