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
            Method: "AddNews",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
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

    beego.GlobalControllerRouter["Vue-cms-server/controllers:SwipeController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:SwipeController"],
        beego.ControllerComments{
            Method: "GetAllSwipe",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers:SwipeController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers:SwipeController"],
        beego.ControllerComments{
            Method: "AddSwipe",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
