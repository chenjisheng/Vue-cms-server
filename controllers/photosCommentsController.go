package controllers

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type PhotosCommentsController struct {
	beego.Controller
}

//@Title /photos/comments/:id
//@Description 根据图片 id 获取图片详情
//@Param id path int true "图片id"
//@router /comments/:id [get]
func (this *PhotosCommentsController) GetComments() {
	o := orm.NewOrm()
	id, err := this.GetInt(":id")
	pageIndex, err := this.GetInt("pageindex")
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	var comments []models.PhotosComments
	o.QueryTable("PhotosComments").OrderBy("-AddTime").Filter("PhotosList", id).Limit(3, 3*(pageIndex-1)).All(&comments)
	this._return(lib.SuccessCode, "", comments)
}

//@Title /photos/comments/:id
//@Description 根据图片 id 添加评论 JSON format:[normal]
//@Param comment body string true "评论内容"
//@Param add_time body string true "添加评论时间[%Y-%m-%d %H:%M:%S]"
//@router /comments/:id [post]
func (this *PhotosCommentsController) AddComment() {
	o := orm.NewOrm()
	comment := models.PhotosComments{}
	var photo = models.PhotosList{}
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	photo.Id = id
	o.Read(&photo)
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &comment)
	comment.PhotosList = &photo
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	logs.Info(comment)
	o.Insert(&comment)
	this._return(lib.SuccessCode, "add img comment success", "")
}

func (this *PhotosCommentsController) _return(code int, message string, data interface{}) {
	var _data = make(lib.NewStringMap)
	_data.Set("code",code)
	_data.Set("message",message)
	_data.Set("data",data)
	this.Data["json"] = _data
	this.ServeJSON()
	return
}
