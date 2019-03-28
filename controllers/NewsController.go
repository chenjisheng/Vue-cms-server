package controllers

import (
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type NewsController struct {
	beego.Controller
	BaseController
}

//@Title Get newsList
//@router / [get]
func (this *NewsController) GetNewsList() {
	defer this.ServeJSON()
	msg := "get news list "
	datas := this.BaseJSON()
	o := orm.NewOrm()
	var newsList = []models.NewsList{}
	_, err := o.QueryTable("NewsList").RelatedSel().All(&newsList)
	if err != nil {
		datas["code"] = 10001
		datas["message"] = msg + "failed"
		this.Data["json"] = datas
		return
	}
	datas["code"] = 0
	datas["message"] = msg + "success"
	datas["data"] = newsList
	this.Data["json"] = datas
	return
}

//@Title Add news
//@router / [post]
func (this *NewsController) AddNews() {
	defer this.ServeJSON()
	datas := this.BaseJSON()
	o := orm.NewOrm()
	var news = models.NewsList{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &news)
	logs.Info(news)
	o.Insert(news.Content)
	_, err = o.Insert(&news)
	if err != nil {
		logs.Info(err)
		datas["code"] = 10001
		datas["message"] = "Add news failed"
		datas["data"] = ""
		this.Data["json"] = datas
		return
	}
	datas["code"] = 0
	datas["message"] = "Add news success"
	datas["data"] = ""
	this.Data["json"] = datas
}

//@Title Get news info
//@router /info/:id [get]
func (this *NewsController) GetNewsInfo() {
	defer this.ServeJSON()
	var info = models.NewsList{}

	msg := "get news info "
	datas := this.BaseJSON()
	id, err := this.GetInt(":id")
	if err != nil {
		datas["code"] = 10001
		datas["message"] = msg + "failed"
		this.Data["json"] = datas
		return
	}
	logs.Info("要查询的详情id为:", id)
	o := orm.NewOrm()
	err = o.QueryTable("NewsList").RelatedSel().Filter("Id", id).One(&info)
	if err != nil {
		datas["code"] = 10001
		datas["message"] = msg + "failed"
		this.Data["json"] = datas
		return
	}
	info.Click += 1
	o.Update(&info)
	datas["code"] = 0
	datas["message"] = msg + "success"
	datas["data"] = info
	this.Data["json"] = datas
	return
}
