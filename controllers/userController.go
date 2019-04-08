package controllers

import (
	"Vue-cms-server/lib"
	"Vue-cms-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserController struct {
	beego.Controller
}

//@Title /user/login
//@Description 后台管理登陆
//@Param username body string true "用户名"
//@Param password body string true "密码"
//@router /login [post]
func (this *UserController) Login() {
	var user models.User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		this._return(lib.GeneralErrorCode, "param error", "")
		return
	}
	passwd := user.Password
	o := orm.NewOrm()
	err = o.Read(&user, "Username")
	if err != nil {
		logs.Error("user not found")
		this._return(lib.GeneralErrorCode, "user not found", "")
		return
	}
	if !lib.CheckPassword(user.Salt, user.Password, passwd) {
		logs.Error("password invalid")
		this._return(lib.GeneralErrorCode, "user or passwd invalid", "")
		return
	}
	secret := beego.AppConfig.String("jwtsecret")
	var info = map[string]interface{}{
		"id": user.Id,
		"token":    user.Token,
		"username": user.Username,
		"nickname": user.Nickname,
	}
	if _, ok := lib.CheckToken(secret, user.Token); ok {
		// token 有效
		logs.Info("token valid")
		this._return(lib.SuccessCode, "", info)
	} else {
		// token 无效
		logs.Info("token expired,generate new token")
		claim := make(jwt.MapClaims)
		claim["username"] = user.Username
		claim["exp"] = time.Now().Add(time.Hour * 24).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			logs.Error("generate token failed")
			this._return(lib.GeneralErrorCode, "generate token failed", "")
			return
		}
		user.Token = tokenString
		info["token"] = tokenString
		this._return(lib.SuccessCode, "", info)
	}
	user.LastLogin = time.Now()
	o.Update(&user)
	this.Ctx.SetCookie("username", user.Username)
	return
}

//@Title /user/register
//@Description 后台注册用户
//@Param username body string true "用户名"
//@Param password body string true "密码"
//@Param nickname body string true "别名"
//@router /register [post]
func (this *UserController) Register(){
	var user models.User
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&user)
	if err != nil {
		logs.Error("param error")
		this._return(lib.GeneralErrorCode,"param error","")
		return
	}
	if user.Username == "" || user.Password == "" {
		logs.Error("name or password empty")
		this._return(lib.GeneralErrorCode,"name or password empty","")
		return
	}
	logs.Info("begin create user ")
	o := orm.NewOrm()
	err = o.Read(&user,"Username")
	if err == nil {
		logs.Error("user exist")
		this._return(lib.GeneralErrorCode,"user exist","")
		return
	}
	salt := lib.GenerateSalt(12)
	passwd := lib.EncryptStr(salt,user.Password)
	if user.Nickname == "" {
		user.Nickname = user.Username
	}
	user.Password = passwd
	user.Salt = salt
	user.Role = "general"
	o.Insert(&user)
	info := map[string]interface{}{
		"id":user.Id,
		"username":user.Username,
		"nickname":user.Nickname,
	}
	this._return(lib.SuccessCode,"add user success",info)
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
