package models

type GoodsSwipe struct {
	Id int `json:"id"`
	Imgurl string `json:"img_url"`
	GoodsList *GoodsList `json:"-" orm:"rel(fk)"`
}
