// @APIVersion 1.0.0
// @Title Vue CMS API
// @Description "Vue CMS API 服务"
// @Contact mail_maomao@613.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Vue-cms-server/controllers"
	"Vue-cms-server/controllers/admin"
	"Vue-cms-server/lib"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
)

var secret = beego.AppConfig.String("jwtsecret")

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorizatioin", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/v1/admin/*", beego.BeforeRouter, FileterFunc)
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/swipe",
			beego.NSInclude(
				&controllers.SwipeController{},
			),
		),
		// 新闻类路由
		beego.NSNamespace("/news",
			beego.NSInclude(
				&controllers.NewsController{},
			),
			beego.NSInclude(
				&controllers.CommentsController{}),
		),
		// 图片类路由
		beego.NSNamespace("/photos",
			beego.NSInclude(
				&controllers.PhotosController{},
			),
			beego.NSInclude(
				&controllers.PhotosCommentsController{},
			),
		),
		// 商品类路由
		beego.NSNamespace("/goods",
			beego.NSInclude(
				&controllers.GoodsController{},
			),
		),
		beego.NSNamespace("/video",
			beego.NSInclude(
				&controllers.VideoController{}),
		),
		// 登陆注册路由
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{}),
		),
		// 需要登陆认证路由
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&admin.SwipeController{},
				&admin.UserController{},
				&admin.GoodsController{},
				&admin.NewsController{},
				&admin.PhotosController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

var FileterFunc = func(ctx *context.Context) {
	// 判断是否有cookie
	userName := ctx.Input.Cookie("username")
	logs.Info("当前cookie代表用户为:", userName)
	info := map[string]interface{}{
		"code":    lib.GeneralErrorCode,
		"message": "",
		"data":    "",
	}
	if userName == "" {
		logs.Error("no cookies")
		info["message"] = "no cookies"
		ctx.Output.JSON(info, true, true)
		return
	}
	authString := ctx.Input.Header("Authorization")
	username, ok := lib.CheckToken(secret, authString)
	if !ok {
		logs.Error("token invalid")
		info["message"] = "token invalid"
		ctx.Output.JSON(info, true, true)
		return
	}
	if userName != username {
		logs.Error("token not match cookie")
		info["message"] = "token or cookie invalid"
		ctx.Output.JSON(info, true, true)
		return
	}
	logs.Info("token and cookie check pass")
	return
}
