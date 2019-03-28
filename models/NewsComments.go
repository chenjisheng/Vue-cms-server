package models

import "Vue-cms-server/lib"

type NewsCommnets struct {
	Id int `json:"id"`
	Comment string `json:"comment"`
	AddTime lib.Time `json:"add_time"`
	NewsList *NewsList `orm:"rel(fk)" json:"-"`
}
