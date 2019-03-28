package controllers

import (
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type SwipeController struct {
	beego.Controller
	BaseController
}

//@Title Get all swipe list
//@router / [get]
func (this *SwipeController) GetAllSwipe(){
	defer this.ServeJSON()
	o := orm.NewOrm()
	datas := this.BaseJSON()
	swipeList := []models.Swipe{}
	o.QueryTable("Swipe").All(&swipeList)
	datas["code"] = 0
	datas["data"] = swipeList
	this.Data["json"] = datas
	return
}

//@Title Add swipe
//@router / [post]
func (this *SwipeController) AddSwipe(){
	defer this.ServeJSON()
	datas := this.BaseJSON()
	o := orm.NewOrm()
	var swipe = []models.Swipe{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&swipe)
	if err != nil {
		msg := "Add swipe params failed"
		logs.Info(msg)
		datas["code"] = 10001
		datas["message"] = msg
		this.Data["json"] = datas
		return
	}
	id,err := o.InsertMulti(len(swipe),swipe)
	if err != nil {
		msg := "Add swipe params failed"
		logs.Info(msg,err)
		datas["code"] = 10001
		datas["message"] = msg
		this.Data["json"] = datas
		return
	}
	msg := "Add swipe success"
	logs.Info(msg)
	datas["code"] = 0
	datas["message"] = msg
	datas["data"] = id
	this.Data["json"] = datas
	return
}