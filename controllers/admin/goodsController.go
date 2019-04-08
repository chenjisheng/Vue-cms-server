package admin

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsController struct {
	beego.Controller
}
//@Title /admin/goods/add
//@Description 添加商品
//@Param title body string true "商品标题"
//@Param add_time body string true "添加时间"
//@Param content body string true "商品描述"
//@Param click body int true "点击次数"
//@Param img_url body string true "商品图片"
//@Param sell_price body int true "售卖价格"
//@Param market_price body int true "标记价格"
//@Param stock_quantity body int true "库存"
//@router /goods/add [post]
func (this *GoodsController) AddGoods(){
	var goods models.GoodsList
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&goods)
	if err != nil {
		this._return(lib.GeneralErrorCode,"params error","")
		return
	}
	o := orm.NewOrm()
	id,_ := o.Insert(&goods)
	this._return(lib.SuccessCode,"添加商品成功",id)
}
//@Title /admin/goods/swipe/:id
//@Description 根据 商品列表id 添加轮播图[json:切片形式]
//@Param id path int true "商品列表id"
//@Param img_url body string true "轮播图地址"
//@router /goods/swipe/:id [post]
func (this *GoodsController) AddGoodsSwipe(){
	goodsId ,err := this.GetInt(":id")
	if err != nil {
		logs.Error("获取商品id 失败")
		this._return(lib.GeneralErrorCode,"param error","")
		return
	}
	var swipe []models.GoodsSwipe
	err = json.Unmarshal(this.Ctx.Input.RequestBody,&swipe)
	if err != nil {
		logs.Error("param error")
		this._return(lib.GeneralErrorCode,"param error","")
		return
	}
	o := orm.NewOrm()
	var goodsList models.GoodsList
	goodsList.Id = goodsId
	err = o.Read(&goodsList,"Id")
	if err != nil {
		logs.Error("无此商品")
		this._return(lib.GeneralErrorCode,"goods not found","")
		return
	}
	for k,_ := range swipe{
		swipe[k].GoodsList = &goodsList
	}
	o.InsertMulti(len(swipe),swipe)
	this._return(lib.SuccessCode,"add goods swipe success",len(swipe))
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
