package models
// 图片详情页面缩略图表
type PhotosHum struct {
	Id int `json:"id"`
	HumUrl string `json:"src"`
	PhotoList *PhotosList `orm:"rel(fk)" json:"-"`
}
