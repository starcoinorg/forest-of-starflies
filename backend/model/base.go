package model

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func DBInit() {
	dbhost, _ := beego.AppConfig.String("dbhost")
	dbport, _ := beego.AppConfig.String("dbport")
	dbuser, _ := beego.AppConfig.String("dbuser")
	dbpassword, _ := beego.AppConfig.String("dbpassword")
	dbname, _ := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dbSource := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dbSource)
	orm.RegisterModel(new(Peer))
}
