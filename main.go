package main

import (
	"peers_crawler/crawler"
	"peers_crawler/model"

	"github.com/astaxie/beego/toolbox"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	model.DBInit()

	GetMainPeers := toolbox.NewTask("GetMainPeers", "0 */5 * * * *", crawler.GetMainPeers)
	GetBarnardPeers := toolbox.NewTask("GetBarnardPeers", "0 */5 * * * *", crawler.GetBarnardPeers)
	GetProximaPeers := toolbox.NewTask("GetProximaPeers", "0 */5 * * * *", crawler.GetProximaPeers)
	toolbox.AddTask("GetMainPeers", GetMainPeers)
	toolbox.AddTask("GetBarnardPeers", GetBarnardPeers)
	toolbox.AddTask("GetProximaPeers", GetProximaPeers)

	toolbox.StartTask()

}

func main() {
	orm.RunSyncdb("default", false, true)
	web.Run()
}
