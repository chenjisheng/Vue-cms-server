package models

import (
	"Vue-cms-server/lib"
)
// 新闻列表
type NewsList struct {
	Id int `json:"id"`
	Title string `json:"title"`
	AddTime lib.Time `json:"add_time";orm:"type(date);auto_now_add"`
	Click int `json:"click"`
	Url string `json:"url"`
	Content *NewsContents `orm:"rel(one);" json:"content"`
	Comments []*NewsComments `orm:"reverse(many)" json:"-"`
}

