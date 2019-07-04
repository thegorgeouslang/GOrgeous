// Author: James Mallon <jamesmallondev@gmail.com>
// configs package - the goal of the file is having a singleton form of reading the
// .env file, not reopening it time after time - while it makes necessary to
// recompile the code after adding new values to the .env file, it improves the
// performance
package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// environment values
var Env map[string]string

// init function - data and process initialization
func init() {
	envFile := os.Getenv("GOPATH") + "/src/GoAuthorization/.env"
	e := godotenv.Load(envFile)
	if e != nil {
		//var lg *log.Logger
		//file, _ := os.OpenFile(os.Getenv("GOPATH")+"/src/GoAuthorization/logs/config.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 1444)
		//lg = log.New(file, "", log.LstdFlags|log.Lshortfile)
		//lg.Printf("Error loading the .env file - %s", e.Error())
		log.Fatal(e.Error())
	}
	Env, e = godotenv.Read()
}
