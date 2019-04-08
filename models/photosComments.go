package models

import "Vue-cms-server/lib"
// 图片评论表
type PhotosComments struct {
	Id int `json:"id"`
	Comment string `json:"comment"`
	AddTime lib.Time `json:"add_time"`
	PhotosList *PhotosList `orm:"rel(fk)" json:"-"`
}
