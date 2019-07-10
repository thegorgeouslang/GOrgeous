// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package layout

import (
	log "TheGorgeous/libs/logger"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"
)

// Struct type flashMessenger -
type FlashMessenger struct{}

//  Set function - it creates the cookie and set the flash message
func (this *FlashMessenger) Set(w *http.ResponseWriter, msg map[string]string) {
	jsMsg, _ := json.Marshal(msg)

	c := http.Cookie{ // creates the cookie
		Name:  "flashmessenger",
		Value: base64.URLEncoding.EncodeToString(jsMsg),
	}
	http.SetCookie(*w, &c)
}

// Get method - it prints the message stored in the cookie
func (this *FlashMessenger) Get(w http.ResponseWriter, r *http.Request) (flshMsg *map[string]string) {
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
		val, e := base64.URLEncoding.DecodeString(c.Value) // decode the value of the cookie

		_ = json.Unmarshal([]byte(val), &flshMsg)
		http.SetCookie(w, &rc) // override the values of the cookie setted in the setMessage function
		if e != nil {
			log.Write("Error", e.Error(), log.Trace())
			return
		}
		return
	}
}
