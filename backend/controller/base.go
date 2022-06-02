package controller

import "github.com/beego/beego/v2/server/web"

// Result common result
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// BaseController just base controller
type BaseController struct {
	web.Controller
}

// ReturnBack simplify the return result
func (c *BaseController) ReturnBack(r Result, e error) {
	if e != nil {

		if r.Code == 0 {
			r.Code = -1
		}
		if r.Msg == "" {
			r.Msg = "error"
		}

	} else {
		if r.Msg == "" {
			r.Msg = "success"
		}
	}
	c.Data["json"] = r
	c.ServeJSON()
}
