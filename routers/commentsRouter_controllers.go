package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Vue-cms-server/controllers:CommentsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:CommentsController"],
        beego.ControllerComments{
            Method: "GetComments",
            Router: `/comments/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:CommentsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:CommentsController"],
        beego.ControllerComments{
            Method: "AddComment",
            Router: `/comments/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:GoodsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetGoodsInfo",
            Router: `/info/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:GoodsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetAllList",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:GoodsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetGoodsSwipe",
            Router: `/swipe/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:NewsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:NewsController"],
        beego.ControllerComments{
            Method: "GetNewsList",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:NewsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:NewsController"],
        beego.ControllerComments{
            Method: "GetNewsInfo",
            Router: `/info/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosCommentsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosCommentsController"],
        beego.ControllerComments{
            Method: "GetComments",
            Router: `/comments/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosCommentsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosCommentsController"],
        beego.ControllerComments{
            Method: "AddComment",
            Router: `/comments/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosController"],
        beego.ControllerComments{
            Method: "GetHumInfo",
            Router: `/humInfo/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosController"],
        beego.ControllerComments{
            Method: "GetPhotosInfo",
            Router: `/info/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosController"],
        beego.ControllerComments{
            Method: "GetImgList",
            Router: `/list/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:PhotosController"],
        beego.ControllerComments{
            Method: "GetPhotoType",
            Router: `/types`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:SwipeController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:SwipeController"],
        beego.ControllerComments{
            Method: "GetAllSwipe",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:UserController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:UserController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:VideoController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:VideoController"],
        beego.ControllerComments{
            Method: "GetAllVideo",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
