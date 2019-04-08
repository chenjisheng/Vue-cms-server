package controllers

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsController struct {
	beego.Controller
}
//@Title /goods/list
//@Description 获取所有端口列表
//@Param pageindex query int true "当前页数"
//@router /list [get]
func (this *GoodsController) GetAllList() {
	goodsList := []models.GoodsList{}
	pageIndex,err := this.GetInt("pageindex")
	if err != nil {
		this._return(lib.GeneralErrorCode,"param error","")
		return
	}
	o := orm.NewOrm()
	o.QueryTable("GoodsList").Limit(3,3*(pageIndex-1)).All(&goodsList)
	this._return(lib.SuccessCode,"",goodsList)
}
//@Title /goods/swipe/:id
//@Description 根据商品id 获取轮播图列表
//@Param id path int true "商品id"
//@router /swipe/:id [get]
func (this *GoodsController) GetGoodsSwipe(){
	id,err := this.GetInt(":id")
	if err != nil {
		logs.Error("获取商品id 失败")
		this._return(lib.GeneralErrorCode,"param error","")
		return
	}
	var swipe []models.GoodsSwipe
	o := orm.NewOrm()
	o.QueryTable("GoodsSwipe").Filter("GoodsList",id).All(&swipe)
	this._return(lib.SuccessCode,"get goods swipe success",swipe)
	return
}

//@Title /goods/info/:id
//@Description 根据商品id 获取轮播图列表
//@Param id path int true "商品id"
//@router /info/:id [get]
func (this *GoodsController) GetGoodsInfo() {
	id,err := this.GetInt(":id")
	if err != nil {
		logs.Error("获取商品id失败")
		this._return(lib.GeneralErrorCode,"param error","")
		return
	}
	var info models.GoodsList
	o := orm.NewOrm()
	info.Id = id
	err = o.Read(&info)
	if err != nil {
		logs.Error("无商品信息")
		this._return(lib.GeneralErrorCode,"param error","")
		return
	}
	this._return(lib.SuccessCode,"get goods info success",info)
	return
}

func (this *GoodsController) _return(code int, message string, data interface{}) {
	var _data =  make(lib.NewStringMap)
	_data.Set("code",code)
	_data.Set("message",message)
	_data.Set("data",data)
	this.Data["json"] = _data
	this.ServeJSON()
	return
}
