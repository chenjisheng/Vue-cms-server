package admin

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

//@Title /admin/update/:id
//@Description 后台注册用户信息更新,不允许更新用户名
//@Param username body string false "用户名"
//@Param password body string false "密码"
//@Param nickname body string false "别名"
//@Param id path int true "用户id"
//@router /update/:id [post]
func (this *UserController) UpdateUser(){
	var inputUser models.User
	var oldUser models.User
	id,err := this.GetInt(":id")
	if err != nil {
		logs.Error("url error")
		this._return(lib.GeneralErrorCode,"url error","")
		return
	}
	err = json.Unmarshal(this.Ctx.Input.RequestBody,&inputUser)
	if err != nil {
		logs.Error("param error")
		this._return(lib.GeneralErrorCode,"param error","")
		return
	}
	o := orm.NewOrm()
	oldUser.Id = id
	err = o.Read(&oldUser,"Id")
	if err != nil {
		logs.Error("user not found")
		this._return(lib.GeneralErrorCode,"user not found","")
		return
	}
	logs.Info("begin update user ")
	// 普通用户只能更新自己的信息
	if username,ok := this.checkAdminPermission();!ok {
		logs.Info("当前为普通用户更新模式")
		if username != oldUser.Username {
			logs.Error("no permission")
			this._return(lib.GeneralErrorCode,"no permission","")
			return
		}
		if inputUser.Nickname != ""{
			oldUser.Nickname = inputUser.Nickname
		}
		if inputUser.Password != "" {
			oldUser.Password = lib.EncryptStr(oldUser.Salt,inputUser.Password)
		}
		if inputUser.Role != "" {
			logs.Error("no permission")
			this._return(lib.GeneralErrorCode,"no permission","")
			return
		}
	} else {
	// 超级管理员可以更改所有用户信息
		logs.Info("超级管理员模式")
		if inputUser.Password != "" {
			oldUser.Password = lib.EncryptStr(oldUser.Salt,inputUser.Password)
		}
		if inputUser.Username == username  {
			if inputUser.Nickname != "" {
				logs.Error("do not change super user ")
				this._return(lib.GeneralErrorCode,"do not change super user","")
				return
			}
			if inputUser.Role != "" {
				logs.Error("do not change super user ")
				this._return(lib.GeneralErrorCode,"do not change super user","")
				return
			}
		} else {
			logs.Error("do not change super user")
			this._return(lib.GeneralErrorCode,"do not change super user","")
			return
		}
		if inputUser.Role != "" {
				if inputUser.Role == "general" || inputUser.Role == "admin" {
					oldUser.Role = inputUser.Role
				} else {
					logs.Error("role error")
					this._return(lib.GeneralErrorCode,"role error","")
					return
				}
		}
		if inputUser.Nickname != "" {
			oldUser.Nickname = inputUser.Nickname
		}
	}
	_,err = o.Update(&oldUser)
	if err != nil {
		logs.Error("update user: ",oldUser.Username,"info failed")
		this._return(lib.GeneralErrorCode,"update user info falied","")
		return
	}
	logs.Info("update user",oldUser.Username,"info success")
	this._return(lib.SuccessCode,"update user info success",oldUser.Id)
}
func (this *UserController) _return(code int, message string, data interface{}) {
	var _data = lib.NewStringMap{}
	_data.Set("code", code)
	_data.Set("message", message)
	_data.Set("data", data)
	this.Data["json"] = _data
	this.ServeJSON()
	return
}
func (this *UserController) checkAdminPermission() (loginUser string,res bool) {
	o := orm.NewOrm()
	loginUser = this.Ctx.GetCookie("username")
	user := models.User{Username:loginUser}
	err := o.Read(&user,"Username")
	if err != nil {
		logs.Error("查询用户错误")
		return loginUser,false
	}
	if user.Role != "administrator" {
		logs.Error("无权限")
		return loginUser,false
	}
	return loginUser,true
}