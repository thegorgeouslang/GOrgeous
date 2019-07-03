// Author: James Mallon <jamesmallondev@gmail.com>
// databases package -
package databases

import (
	"github.com/jinzhu/gorm"
)

type DBInterface interface {
	//Insert(obj interface{}) error
	//Select(obj interface{}) error
	//SelectOne(obj interface{}) error
	//Update(obj interface{}, changes map[string]interface{}) error
	//Delete(obj interface{}) error
	GetConn() *gorm.DB
	Check() error
	Migrate(ojb ...interface{})
}
