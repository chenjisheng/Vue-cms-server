package models

import "Vue-cms-server/lib"

type PhotosList struct {
	Id int `json:"id"`
	Title string `json:"title"`
	AddTime lib.Time `json:"add_time"`
	Click int `json:"click"`
	ImgUrl string `json:"img_url"`
	Content string `json:"content"`
	PhotosInfo *PhotosInfo `orm:"rel(one)" json:"photo_info"`
	PhotoType *PhotosType `orm:"rel(fk)" json:"-"`
	Comments []*PhotosComments `orm:"reverse(many)" json:"-"`
	PhotosHum []*PhotosHum `orm:"reverse(many)" json:"photo_hum_list"`
}
