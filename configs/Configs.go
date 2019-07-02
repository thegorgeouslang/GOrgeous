// Author: James Mallon <jamesmallondev@gmail.com>
// configs package - the goal of the file is having a singleton form of reading the
// .env file, not reopening it time after time - while it makes necessary to
// recompile the code after adding new values to the .env file, it improves the
// performance
package configs

import (
	"github.com/joho/godotenv"
	"log"
)

// environment values
var Env map[string]string

// init function - data and process initialization
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Env, err = godotenv.Read()
}
