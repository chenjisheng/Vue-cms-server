package controllers

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type NewsController struct {
	beego.Controller
}

//@Title /news/
//@Description 获取新闻列表
//@router / [get]
func (this *NewsController) GetNewsList() {
	msg := "get news list "
	o := orm.NewOrm()
	var newsList = []models.NewsList{}
	_, err := o.QueryTable("NewsList").RelatedSel().All(&newsList)
	if err != nil {
		this._return(lib.GeneralErrorCode, msg+"failed", "")
		return
	}
	this._return(lib.SuccessCode, msg+"success", newsList)
}


//@Title /news/info/:id
//@Description 获取新闻详细内容
//@Param id path int true "新闻列表的id"
//@router /info/:id [get]
func (this *NewsController) GetNewsInfo() {
	var info = models.NewsList{}
	msg := "get news info "
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(lib.GeneralErrorCode, msg+"failed", "")
		return
	}
	logs.Info("要查询的详情id为:", id)
	o := orm.NewOrm()
	err = o.QueryTable("NewsList").RelatedSel().Filter("Id", id).One(&info)
	if err != nil {
		this._return(lib.SuccessCode, msg+"failed", "")
		return
	}
	info.Click += 1
	o.Update(&info)
	this._return(lib.SuccessCode, msg+"success", info)
}

func (this *NewsController) _return(code int, message string, data interface{}) {
	var _data = lib.NewStringMap{}
	_data.Set("code",code)
	_data.Set("message",message)
	_data.Set("data",data)
	this.Data["json"] = _data
	this.ServeJSON()
	return
}
