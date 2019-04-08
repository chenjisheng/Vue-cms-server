package admin

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type PhotosController struct {
	beego.Controller
}

//@Title /admin/photos/list
//@Description 添加图片列表 JSON format:[normal]
//@Param title body string true "图片类型"
//@Param data body models.PhotosList true "图片详情 JSON format:[切片形式]"
//@router /photos/list [post]
func (this *PhotosController) AddImgList() {
	_type := models.PhotosType{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_type)
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	o := orm.NewOrm()
	err = o.Read(&_type, "Type")
	if err != nil {
		this._return(lib.GeneralErrorCode, "no image type", "")
		return
	}
	imgList := struct {
		Data []models.PhotosList
	}{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &imgList)
	if err != nil {
		logs.Info("json format error", err.Error())
		this._return(lib.GeneralErrorCode, "data params error", "")
		return
	}
	if ok := imgList.Data[0].Content; ok == "" {
		logs.Info("no content")
		this._return(lib.GeneralErrorCode, "data params error", "")
		return
	}
	for k, _ := range imgList.Data {
		imgList.Data[k].PhotoType = &_type
		o.Insert(imgList.Data[k].PhotosInfo)
		logs.Info(imgList.Data[k].PhotoType)
	}
	o.InsertMulti(len(imgList.Data), imgList.Data)
	this._return(lib.SuccessCode, "add image list success", "")
}

//@Title /admin/photos/types
//@Description 添加图片类型 JSON format:[切片形式]
//@Param title body string true "图片类型"
//@router /photos/types [post]
func (this *PhotosController) AddPhotoType() {
	photoTypes := []models.PhotosType{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &photoTypes)
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	o := orm.NewOrm()
	_id, err := o.InsertMulti(len(photoTypes), photoTypes)
	if err != nil {
		this._return(lib.GeneralErrorCode, err.Error(), "")
		return
	}
	logs.Info(photoTypes)
	this._return(lib.SuccessCode, "add phototype success", _id)
}
//@Title /admin/photos/humInfo/:id
//@Description 根据图片id 添加图片缩略图列表 JSON format:[切片]
//@Param src body string true "图片地址"
//@Param id path int true "图片列表id"
//@router /photos/humInfo/:id [post]
func (this *PhotosController) AddHumInfo() {
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(lib.GeneralErrorCode, "params error", "")
		return
	}
	huminfo := []models.PhotosHum{}
	photo := models.PhotosList{}
	o := orm.NewOrm()
	photo.Id = id
	o.Read(&photo)
	err = json.Unmarshal(this.Ctx.Input.RequestBody,&huminfo)
	if err != nil {
		this._return(lib.GeneralErrorCode, "json params error", "")
		return
	}
	for k := range huminfo {
		huminfo[k].PhotoList = &photo
	}
	o.InsertMulti(len(huminfo),huminfo)
	this._return(lib.SuccessCode, "add photos hum success", "")
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
