// Author: James Mallon <jamesmallondev@gmail.com>
// databases package -
package databases

import (
	conf "GoAuthorization/configs"
	log "GoAuthorization/libs/logger"
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
		log.Write("error", e.Error(), log.Trace())
		panic(e)
	}
	DbConn = &dbConn{db: db}
}

// check method -
func (this *dbConn) check() (e error) {
	if faults := this.db.GetErrors(); len(faults) > 0 {
		e = faults[0]
	}
	return
}

// Insert method -
func (this *dbConn) Insert(obj interface{}) (e error) {
	this.db.Create(obj)
	return this.check()
}

// Select method -
func (this *dbConn) Select(obj interface{}) (e error) {
	return e
}

// SelectOne method -
func (this *dbConn) SelectOne(obj interface{}) (e error) {
	this.db.Take(obj)
	return this.check()
}

// Update method -
func (this *dbConn) Update(obj interface{}) (e error) {
	return e
}

// Delete method -
func (this *dbConn) Delete(obj interface{}) (e error) {
	return e
}

//  Migrate method - create tables based on models
func (this *dbConn) Migrate(models ...interface{}) {
	this.db.Debug().AutoMigrate(models...)
}
