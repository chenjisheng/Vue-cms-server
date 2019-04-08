package controllers

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type PhotosController struct {
	beego.Controller
}


//@Title /photos/types
//@Description 获取图片类型列表
//@router /types [get]
func (this *PhotosController) GetPhotoType() {
	o := orm.NewOrm()
	types := []models.PhotosType{}
	o.QueryTable("PhotosType").All(&types)
	this._return(lib.SuccessCode, "get photo types success", types)
}


//@Title /photos/humInfo/:id
//@Description 添加图片列表 JSON format:[切片]
//@Param id path int true "图片列表id"
//@router /humInfo/:id [get]
func (this *PhotosController) GetHumInfo(){
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	humlist := []models.PhotosHum{}
	o := orm.NewOrm()
	o.QueryTable("PhotosHum").Filter("PhotoList",id).All(&humlist)
	this._return(lib.SuccessCode,"get hum list success",humlist)
}
//@Title /list/:id
//@Description 根据图片类型获取图片列表,0 默认为全部列表
//@Param id path int true "图片类型的id"
//@router /list/:id [get]
func (this *PhotosController) GetImgList() {
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	var imgList []models.PhotosList
	o := orm.NewOrm()
	if id == 0 {
		o.QueryTable("PhotosList").RelatedSel().All(&imgList)
	} else {
		o.QueryTable("PhotosList").RelatedSel().Filter("PhotoType", id).All(&imgList)
	}
	this._return(lib.SuccessCode, "get image list success", imgList)
}

//@Title /info/:id
//@Description 根据图片id 查找图片详细信息
//@Param id path int true "图片的id"
//@router /info/:id [get]
func (this *PhotosController) GetPhotosInfo() {
	o := orm.NewOrm()
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	info := models.PhotosList{}
	o.QueryTable("PhotosList").RelatedSel().Filter("PhotosInfo", id).One(&info)
	info.Click += 1
	o.Update(&info)
	logs.Info(info.PhotosHum)
	this._return(lib.SuccessCode, "get img info success", info)
}
func (this *PhotosController) _return(code int, message string, data interface{}) {
	var _data = make(lib.NewStringMap)
	_data.Set("code",code)
	_data.Set("message",message)
	_data.Set("data",data)
	this.Data["json"] = _data
	this.ServeJSON()
	return
}
