// Author: James Mallon <jamesmallondev@gmail.com>
// helpers package -
package helpers

import (
	. "TheGorgeous/libs/layout"
)

var Flashmsg = FlashMessenger{}

// init function - data and process initialization
func init() {
	Flashmsg.CkName = "flashmessenger"
}
