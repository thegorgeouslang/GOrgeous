package logger

import (
	conf "TheGorgeous/configs"
	"fmt"
	"github.com/thegorgeouslang/logit"
	"os"
	"time"
)

var Logit = *logit.Syslog

func init() {
	Logit.Filepath = fmt.Sprintf("%s/src/%s/%s%s%s",
		os.Getenv("GOPATH"),
		conf.Env["project_name"],
		conf.Env["logfile_path"],        // logfile path
		time.Now().Format("2006_01_02"), // logfile name
		conf.Env["logfile_ext"])         // logfile extension
}
