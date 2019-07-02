// Author: James Mallon <jamesmallondev@gmail.com>
// databases package -
package databases

type DBInterface interface {
	Insert(obj interface{}) error
	Select(obj interface{}) error
	SelectOne(obj interface{}) error
	Update(obj interface{}) error
	Delete(obj interface{}) error
	Migrate(ojb ...interface{})
}
