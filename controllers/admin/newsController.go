package admin

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type NewsController struct {
	beego.Controller
}

//@Title /admin/news/add
//@Description 添加新闻列表 JSON format:[normal]
//@Param title body string true "新闻标题"
//@Param click body int true "新闻点击次数"
//@Param url body string true "新闻图片缩略图地址"
//@Param add_time body string true "新闻发表时间"
//@Param content body models.NewsContents true "新闻详细内容"
//@router /news/add [post]
func (this *NewsController) AddNews() {
	o := orm.NewOrm()
	var news = models.NewsList{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &news)
	logs.Info(news)
	o.Insert(news.Content)
	_, err = o.Insert(&news)
	if err != nil {
		logs.Info(err)
		this._return(lib.GeneralErrorCode, "Add news failed", "")
		return
	}
	this._return(lib.SuccessCode, "Add news success", "")
}
func (this *NewsController) _return(code int, message string, data interface{}) {
	var _data = lib.NewStringMap{}
	_data.Set("code",code)
	_data.Set("message",message)
	_data.Set("data",data)
	this.Data["json"] = _data
	this.ServeJSON()
	return
}
