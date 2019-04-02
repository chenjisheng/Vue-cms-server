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
		this._return(10003, msg+"failed", "")
	}
	this._return(0, msg+"success", newsList)
}

//@Title /news/
//@Description 添加新闻列表 JSON format:[normal]
//@Param title body string true "新闻标题"
//@Param click body int true "新闻点击次数"
//@Param url body string true "新闻图片缩略图地址"
//@Param add_time body string true "新闻发表时间"
//@Param content body models.NewsContents true "新闻详细内容"
//@router / [post]
func (this *NewsController) AddNews() {
	o := orm.NewOrm()
	var news = models.NewsList{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &news)
	logs.Info(news)
	o.Insert(news.Content)
	_, err = o.Insert(&news)
	if err != nil {
		logs.Info(err)
		this._return(10003, "Add news failed", "")
	}
	this._return(0, "Add news success", "")
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
		this._return(10003, msg+"failed", "")
	}
	logs.Info("要查询的详情id为:", id)
	o := orm.NewOrm()
	err = o.QueryTable("NewsList").RelatedSel().Filter("Id", id).One(&info)
	if err != nil {
		this._return(10003, msg+"failed", "")
	}
	info.Click += 1
	o.Update(&info)
	this._return(0, msg+"success", info)
}

func (this *NewsController) _return(code int, message string, data interface{}) {
	var datas = make(map[string]interface{})
	datas["code"] = code
	datas["message"] = message
	datas["data"] = data
	this.Data["json"] = datas
	this.ServeJSON()
	return
}
