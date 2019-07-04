// Author: James Mallon <jamesmallondev@gmail.com>
// configs package -
package configs

import (
	"testing"
)

// Test function TestInit to evaluate the Init
func TestInit(t *testing.T) {
	if Env["project_name"] != "GoAuthorization" { // check the response for errors
		t.Errorf("The value should be GoAuthorization")
	}
}
