package controllers

import (
	"Vue-cms-server/lib"
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404(){
	url := this.Ctx.Input.URL()
	this._return(lib.ServerNotFount,url + " not found","")
}

func (this *ErrorController) Error500(){
	this._return(lib.ServerErrorCode,"server error","")
}

func (this *ErrorController) _return(code int, message string, data interface{}) {
	var  _data = make(lib.NewStringMap)
	_data.Set("code",code)
	_data.Set("message",message)
	_data.Set("data",data)
	this.Data["json"] = _data
	this.ServeJSON()
	return
}
