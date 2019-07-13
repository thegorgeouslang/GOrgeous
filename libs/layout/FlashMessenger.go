// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package layout

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"
)

// Struct type flashMessenger -
type FlashMessenger struct {
	CkName string
}

//  Set function - it creates the cookie and set the flash message
func (this *FlashMessenger) Set(w *http.ResponseWriter, msg map[string]string) (e error) {
	jsMsg, e := json.Marshal(msg)

	if e != nil {
		return
	}

	c := http.Cookie{ // creates the cookie
		Name:  this.CkName,
		Value: base64.URLEncoding.EncodeToString(jsMsg),
	}
	http.SetCookie(*w, &c)
	return
}

// Get method - it prints the message stored in the cookie
func (this *FlashMessenger) Get(w http.ResponseWriter, r *http.Request) (*map[string]string, error) {
	var flshMsg *map[string]string

	c, e := r.Cookie(this.CkName)
	if e != nil { // check for errors - check if the cookie with the specific name is setted
		return flshMsg, e
	} else {
		rc := http.Cookie{
			Name:    this.CkName,
			MaxAge:  -1,              // a negative number means it's already expired
			Expires: time.Unix(1, 0), // an expiration value that's in the past - already expired
		}
		val, e := base64.URLEncoding.DecodeString(c.Value) // decode the value of the cookie

		if e != nil {
			return flshMsg, e
		}

		_ = json.Unmarshal([]byte(val), &flshMsg)
		http.SetCookie(w, &rc) // override the values of the cookie setted in the setMessage function
		return flshMsg, e

	}
}
