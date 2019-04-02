package models

import "github.com/astaxie/beego/orm"
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
	)
	orm.RunSyncdb("default", false, true)
}
