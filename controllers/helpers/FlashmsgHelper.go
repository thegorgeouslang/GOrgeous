// Author: James Mallon <jamesmallondev@gmail.com>
// helpers package -
package helpers

import (
	. "github.com/thegorgeouslang/FlashMessenger.git"
)

var Flashmsg = FlashMessenger{}

// init function - data and process initialization
func init() {
	Flashmsg.CkName = "flashmessenger"
}
