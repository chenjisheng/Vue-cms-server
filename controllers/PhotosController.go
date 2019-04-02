package controllers

import (
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type PhotosController struct {
	beego.Controller
}

//@Title /photos/types
//@Description 添加图片类型 JSON format:[切片形式]
//@Param title body string true "图片类型"
//@router /types [post]
func (this *PhotosController) AddPhotoType() {
	photoTypes := []models.PhotosType{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &photoTypes)
	if err != nil {
		this._return(10003, "params error", "")
	}
	o := orm.NewOrm()
	_id, err := o.InsertMulti(len(photoTypes), photoTypes)
	if err != nil {
		this._return(10003, err.Error(), "")
	}
	logs.Info(photoTypes)
	this._return(0, "add phototype success", _id)
}

//@Title /photos/types
//@Description 获取图片类型列表
//@router /types [get]
func (this *PhotosController) GetPhotoType() {
	o := orm.NewOrm()
	types := []models.PhotosType{}
	o.QueryTable("PhotosType").All(&types)
	this._return(0, "get photo types success", types)
}

//@Title /photos/list
//@Description 添加图片列表 JSON format:[normal]
//@Param title body string true "图片类型"
//@Param data body models.PhotosList true "图片详情 JSON format:[切片形式]"
//@router /list [post]
func (this *PhotosController) AddImgList() {
	_type := models.PhotosType{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_type)
	if err != nil {
		this._return(10003, "params error", "")
	}
	o := orm.NewOrm()
	err = o.Read(&_type, "Type")
	if err != nil {
		this._return(10003, "no image type", "")
	}
	imgList := struct {
		Data []models.PhotosList
	}{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &imgList)
	if err != nil {
		logs.Info("json format error", err.Error())
		this._return(10003, "data params error", "")
	}
	if ok := imgList.Data[0].Content; ok == "" {
		logs.Info("no content")
		this._return(10003, "data params error", "")
	}
	for k, _ := range imgList.Data {
		imgList.Data[k].PhotoType = &_type
		o.Insert(imgList.Data[k].PhotosInfo)
		logs.Info(imgList.Data[k].PhotoType)
	}
	o.InsertMulti(len(imgList.Data), imgList.Data)
	this._return(0, "add image list success", "")
}

//@Title /photos/humInfo/:id
//@Description 添加图片列表 JSON format:[切片]
//@Param src body string true "图片地址"
//@Param id path int true "图片列表id"
//@router /humInfo/:id [post]
func (this *PhotosController) AddHumInfo() {
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(10003, "params error", "")
	}
	huminfo := []models.PhotosHum{}
	photo := models.PhotosList{}
	o := orm.NewOrm()
	photo.Id = id
	o.Read(&photo)
	err = json.Unmarshal(this.Ctx.Input.RequestBody,&huminfo)
	if err != nil {
		this._return(10003, "json params error", "")
	}
	for k := range huminfo {
		huminfo[k].PhotoList = &photo
	}
	o.InsertMulti(len(huminfo),huminfo)
	this._return(0, "add photos hum success", "")
}

//@Title /photos/humInfo/:id
//@Description 添加图片列表 JSON format:[切片]
//@Param id path int true "图片列表id"
//@router /humInfo/:id [get]
func (this *PhotosController) GetHumInfo(){
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(10003, "params error", "")
	}
	humlist := []models.PhotosHum{}
	o := orm.NewOrm()
	o.QueryTable("PhotosHum").Filter("PhotoList",id).All(&humlist)
	this._return(0,"get hum list success",humlist)
}
//@Title /list/:id
//@Description 根据图片类型获取图片列表,0 默认为全部列表
//@Param id path int true "图片类型的id"
//@router /list/:id [get]
func (this *PhotosController) GetImgList() {
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(10003, "params error", "")
	}
	var imgList []models.PhotosList
	o := orm.NewOrm()
	if id == 0 {
		o.QueryTable("PhotosList").RelatedSel().All(&imgList)
	} else {
		o.QueryTable("PhotosList").RelatedSel().Filter("PhotoType", id).All(&imgList)
	}
	this._return(0, "get image list success", imgList)
}

//@Title /info/:id
//@Description 根据图片id 查找图片详细信息
//@Param id path int true "图片的id"
//@router /info/:id [get]
func (this *PhotosController) GetPhotosInfo() {
	o := orm.NewOrm()
	id, err := this.GetInt(":id")
	if err != nil {
		this._return(10003, "params error", "")
	}
	info := models.PhotosList{}
	o.QueryTable("PhotosList").RelatedSel().Filter("PhotosInfo", id).One(&info)
	info.Click += 1
	o.Update(&info)
	logs.Info(info.PhotosHum)
	this._return(0, "get img info success", info)
}
func (this *PhotosController) _return(code int, message string, data interface{}) {
	var datas = make(map[string]interface{})
	datas["code"] = code
	datas["message"] = message
	datas["data"] = data
	this.Data["json"] = datas
	this.ServeJSON()
	return
}
