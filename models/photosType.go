package models
// 图片类型表
type PhotosType struct {
	Id int `json:"id"`
	Type string `json:"title" orm:"unique"`
	PhotosList []*PhotosList `json:"-" orm:"reverse(many)"`
}

