package controllers

import (
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type CommentsController struct {
	beego.Controller
	BaseController
}


//@Title get comments
//@router /comments/:id [get]
func (this *CommentsController) GetComments(){
	defer this.ServeJSON()
	datas := this.BaseJSON()
	o := orm.NewOrm()
	id,err := this.GetInt(":id")
	pageIndex,err := this.GetInt("pageindex")
	if err != nil {
		datas["code"] = 1
		datas["message"] = "params error"
		this.Data["json"] = datas
		return
	}
	var  comments []models.NewsCommnets
	o.QueryTable("NewsCommnets").Filter("NewsList",id).Limit(3,3*(pageIndex-1)).All(&comments)
	datas["code"] = 0
	datas["data"] = comments
	this.Data["json"] = datas
	return
}

//@Title add comments
//@Param comment body string true "评论内容"
//@router /comments/:id [post]
func (this *CommentsController) AddComment(){
	defer this.ServeJSON()
	datas := this.BaseJSON()
	o := orm.NewOrm()
	comment := models.NewsCommnets{}
	var news = models.NewsList{}
	id,err := this.GetInt(":id")
	if err != nil {
		datas["code"] = 1
		datas["message"] = "params error"
		this.Data["json"] = datas
		return
	}
	news.Id = id
	o.Read(&news)
	err = json.Unmarshal(this.Ctx.Input.RequestBody,&comment)
	comment.NewsList = &news
	if err != nil {
		datas["code"] = 1
		datas["message"] = "params error"
		this.Data["json"] = datas
		return
	}
	logs.Info(comment)
	o.Insert(&comment)
	datas["code"] = 10000
	this.Data["json"] = datas
	return
}
