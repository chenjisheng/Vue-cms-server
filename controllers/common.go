package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) BaseJSON() map[string]interface{}{
	var datas = make(map[string]interface{})
	datas["code"] = 0
	datas["data"] = ""
	datas["message"] = ""
	return datas
}
