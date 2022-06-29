package model

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/spf13/viper"
)

// DBInit init the db
func DBInit() {
	viper.AutomaticEnv()
	dbhost := viper.GetString("DBHOST")
	dbport := viper.GetString("DBPORT")
	dbuser := viper.GetString("DBUSER")
	dbpassword := viper.GetString("DBPASSWORD")
	dbname := viper.GetString("DBNAME")
	if dbport == "" {
		dbport = "3306"
	}
	dbSource := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dbSource)
	orm.RegisterModel(new(Peer))
}
