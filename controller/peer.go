package controller

import (
	"peers_crawler/model"

	"github.com/astaxie/beego/orm"
)

// PeerController controller for peer
type PeerController struct {
	BaseController
}

// Get /peers get
func (c *PeerController) Get() {
	var peers []model.Peer
	o := orm.NewOrm()
	_, err := o.QueryTable(model.Peer{}).All(&peers)
	c.ReturnBack(Result{Data: peers}, err)
}
