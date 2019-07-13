// Author: James Mallon <jamesmallondev@gmail.com>
// databases package -
package databases

import (
	conf "TheGorgeous/configs"
	. "TheGorgeous/libs/logger"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DbConn DBInterface
)

type dbConn struct {
	db *gorm.DB
}

// init function - data and process initialization
// Connection created as a singleton resource along with GORM to add flexeibilty in
// terms of SQL flavor preferences
func init() {
	uri := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		conf.Env["db_host"],
		conf.Env["db_port"],
		conf.Env["db_name"],
		conf.Env["db_user"],
		conf.Env["db_pass"],
	)
	db, e := gorm.Open(conf.Env["db_dbms"], uri)
	if e != nil {
		Logit.WriteLog("error", e.Error(), Logit.GetTraceMsg())
		panic(e)
	}
	DbConn = &dbConn{db: db}
}

// GetConn method -
func (this *dbConn) GetConn() *gorm.DB {
	return this.db
}

// check method -
func (this *dbConn) Check() (e error) {
	if faults := this.db.GetErrors(); len(faults) > 0 {
		e = faults[0]
		Logit.WriteLog("error", e.Error(), Logit.GetTraceMsg())
	}
	return
}

//  Migrate method - create tables based on models
func (this *dbConn) Migrate(models ...interface{}) {
	this.db.Debug().AutoMigrate(models...)
}
