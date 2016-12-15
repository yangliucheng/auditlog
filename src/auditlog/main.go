package main

import (
	"auditlog/models/db"
	_ "auditlog/routers"
	"github.com/astaxie/beego"
)

func main() {
	dataBase := beego.AppConfig.String("dataBase")
	datasource := beego.AppConfig.String("datasource")
	db.InitDB(dataBase, datasource)
	beego.Run()
}
