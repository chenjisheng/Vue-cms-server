package admin

import (
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type SwipeController struct {
	beego.Controller
}

//@Title /admin/swipe/
//@Description 添加轮播图 JSON format:[切片形式]
//@Param img_url body string true "轮播图图片地址"
//@router /swipe [post]
func (this *SwipeController) AddSwipe() {
	o := orm.NewOrm()
	var swipe = []models.Swipe{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &swipe)
	if err != nil {
		msg := "Add swipe params failed"
		logs.Info(msg)
		this._return(10003, msg, "")
		return
	}
	id, err := o.InsertMulti(len(swipe), swipe)
	if err != nil {
		msg := "Add swipe params failed"
		logs.Info(msg, err)
		this._return(10003, msg, "")
		return
	}
	msg := "Add swipe success"
	logs.Info(msg)
	this._return(0, msg, id)
}

func (this *SwipeController) _return(code int, message string, data interface{}) {
	var datas = make(map[string]interface{})
	datas["code"] = code
	datas["message"] = message
	datas["data"] = data
	this.Data["json"] = datas
	this.ServeJSON()
	return
}
