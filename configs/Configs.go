// Author: James Mallon <jamesmallondev@gmail.com>
// configs package - the goal of the file is having a singleton form of reading the
// .env file, not reopening it time after time - while it makes necessary to
// recompile the code after adding new values to the .env file, it improves the
// performance
package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// environment values
var Env map[string]string

// init function - data and process initialization
func init() {
	envFile := os.Getenv("GOPATH") + "/src/GoAuthorization/env.json"
	jsonFile, e := os.Open(envFile)
	if e != nil {
		log.Fatal(e.Error())
	}
	jsonData, e := ioutil.ReadAll(jsonFile)
	if e != nil {
		log.Fatal(e.Error())
	}
	json.Unmarshal(jsonData, &Env)
}
