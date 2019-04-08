package models

import (
	"Vue-cms-server/lib"
	"github.com/astaxie/beego/orm"
)
import _ "github.com/mattn/go-sqlite3"

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./cms.db")
	orm.RegisterModel(new(Swipe),
		new(NewsList),
		new(NewsContents),
		new(NewsComments),
		new(PhotosType),
		new(PhotosList),
		new(PhotosInfo),
		new(PhotosComments),
		new(PhotosHum),
		new(GoodsList),
		new(User),
		new(GoodsSwipe),
	)
	orm.RunSyncdb("default", false, true)
	o := orm.NewOrm()
	err := o.Read(&User{Username:"admin"},"Username")
	if err != nil {
		salt := lib.GenerateSalt(12)
		passwd := lib.EncryptStr(salt,"tellusrd")
		user := User{
			Username:"admin",
			Salt:salt,
			Password:passwd,
			Nickname:"超级管理员",
			Role:"administrator",
		}
		o.Insert(&user)
	}
}
