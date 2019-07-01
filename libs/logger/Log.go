// Author: James Mallon <jamesmallondev@gmail.com>
// logger package -
package logger

import (
	"GoAuthentication/configs"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var lg *log.Logger

func init() {
	// set location of log file
	var logpath = fmt.Sprintf("%s%s%s",
		configs.LOGFILE_PATH,            // logfile path
		time.Now().Format("2006_01_02"), // logfile name
		configs.LOGFILE_EXT)             // logfile extension

	var file, e = os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 1444)

	if e != nil {
		panic(e)
	}
	lg = log.New(file, "", log.LstdFlags)
}

// Write method - writes the message to the log file
func Write(category string, msg string, trace string) {
	lg.Printf("%s: %s on %s", category, msg, trace)
}

// GetTraceMsg method - get the full error stack trace
func Trace() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d PID: %d", frame.File, frame.Line, os.Getpid())
}
