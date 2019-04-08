package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net/http"
	"os"
	"time"
)

type VideoController struct {
	beego.Controller
}

//@Title /video/
//@Description 获取视频列表
//@router / [get]
func (this *VideoController) GetAllVideo(){
	video,err := os.Open("abc.mp4")
	if err != nil {
		logs.Error("打开视频有误,",err.Error())
	}
	defer video.Close()
	http.ServeContent(this.Ctx.ResponseWriter,this.Ctx.Request,"abc.mp4",time.Now(),video)
}
