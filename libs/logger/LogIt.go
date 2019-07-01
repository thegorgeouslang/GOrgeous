// Author: James Mallon <jamesmallondev@gmail.com>
// logger package -
package logger

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/thegorgeouslang/logit"
	"go/build"
	"os"
	"time"
)

//
var It = *logit.Syslog

//
func init() {
	// appending custom categories
	It.AppendCategories(map[string][]string{
		"godotenv": {"Godotenv error:",
			"the godotenv dependency is not working properly"},
	})

	// loading godotenv
	err := godotenv.Load()
	if err != nil {
		// in case of error the path and extension of the log file are being
		// hardcoded
		It.Filepath = fmt.Sprintf("%s/%s%s%s", build.Default.GOPATH, "logs/",
			time.Now().Format("2006_01_02"), ".logs")
		// append
		It.WriteLog("godotenv",
			"the godotenv dependency is not properly working", It.GetTraceMsg())
	} else { // try to get the values of the .env file
		// changing the default log file path
		path := fmt.Sprintf("%s/%s%s%s", build.Default.GOPATH, os.Getenv("logfile_path"),
			time.Now().Format("2006_01_02"), os.Getenv("logfile_ext"),
		)
		It.Filepath = path
	}
}
