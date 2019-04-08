package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "AddGoods",
            Router: `/goods/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "AddGoodsSwipe",
            Router: `/goods/swipe/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:NewsController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:NewsController"],
        beego.ControllerComments{
            Method: "AddNews",
            Router: `/news/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:PhotosController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:PhotosController"],
        beego.ControllerComments{
            Method: "AddHumInfo",
            Router: `/photos/humInfo/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:PhotosController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:PhotosController"],
        beego.ControllerComments{
            Method: "AddImgList",
            Router: `/photos/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:PhotosController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:PhotosController"],
        beego.ControllerComments{
            Method: "AddPhotoType",
            Router: `/photos/types`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:SwipeController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:SwipeController"],
        beego.ControllerComments{
            Method: "AddSwipe",
            Router: `/swipe`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["Vue-cms-server/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/update/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
