package db

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// maximum connections of idle(optional)
	maxIdle int = 30
	// maxxium connections of connected(optional)
	maxConn int = 30
)

/**
 * initialize database
 */
func InitDB(dataBase, dataSource string) {

	// mysql or other database
	switch dataBase {

	case "mysql":
		registerMysql(dataBase, dataSource)
	case "":
		//other database register
	}

	// model register
	orm.RegisterModel(new(AuditLog))
	// param01 alias name of table
	// param02 forced to create table , table create when install our platment
	// param02 whether or not show info when create table,true->not show
	orm.RunSyncdb("default", false, true)
	// test()
}

/**
 * the function declare mysql register
 * @param dataBase : name of database
 * @param dataSource : url of database, place in app.conf
 */
func registerMysql(dataBase, dataSource string) {

	// count of database register
	count := 0
	for {
		// drive register
		err := orm.RegisterDriver(dataBase, orm.DRMySQL)
		if err != nil {
			//	 output err
		}
		// database regsiter
		err = orm.RegisterDataBase("default", dataBase, dataSource, maxIdle, maxConn)
		if err != nil {
			// output err
		}
		count++
		if count == 3 {
			break
		}
	}
}
