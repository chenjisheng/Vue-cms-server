package controllers

import (
	"Vue-cms-server/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SwipeController struct {
	beego.Controller
}

//@Title /swipe/
//@Description 获取首页轮播图列表
//@router / [get]
func (this *SwipeController) GetAllSwipe() {
	o := orm.NewOrm()
	swipeList := []models.Swipe{}
	o.QueryTable("Swipe").All(&swipeList)
	this._return(0, "", swipeList)
}

func (this *SwipeController) _return(code int, message string, data interface{}) {
	var datas = make(map[string]interface{})
	datas["code"] = code
	datas["message"] = message
	datas["data"] = data
	this.Data["json"] = datas
	this.ServeJSON()
	return
}
