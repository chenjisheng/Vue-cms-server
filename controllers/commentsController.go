package controllers

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type CommentsController struct {
	beego.Controller
}

//@Title /news/comments/:id
//@Description 根据新闻 id 获取新闻评论列表
//@Param id path int true "新闻的id"
//@Param pageindex query int true "请求的页面"
//@Success 200 {object} models.Comments
//@router /comments/:id [get]
func (this *CommentsController) GetComments() {
	o := orm.NewOrm()
	id, err := this.GetInt(":id")
	pageIndex, err := this.GetInt("pageindex")
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	var comments []models.NewsComments
	o.QueryTable("NewsComments").OrderBy("-AddTime").Filter("NewsList", id).Limit(3, 3*(pageIndex-1)).All(&comments)
	this._return(lib.SuccessCode, "", comments)
}

//@Title /news/comments/:id
//@Description 根据新闻 id 添加新闻评论 JSON format:[normal]
//@Param comment body string true "评论内容"
//@Param add_time body string true "添加评论时间[%Y-%m-%d %H:%M:%S]"
//@router /comments/:id [post]
func (this *CommentsController) AddComment() {
	o := orm.NewOrm()
	comment := models.NewsComments{}
	var news = models.NewsList{}
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	news.Id = id
	o.Read(&news)
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &comment)
	comment.NewsList = &news
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	logs.Info(comment)
	o.Insert(&comment)
	this._return(lib.SuccessCode, "add news comment success", "")
}

func (this *CommentsController) _return(code int, message string, data interface{}) {
	var _data = lib.NewStringMap{}
	_data.Set("code",code)
	_data.Set("message",message)
	_data.Set("data",data)
	this.Data["json"] = _data
	this.ServeJSON()
	return
}
