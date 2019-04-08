package models

import "time"

type User struct {
	Id int `json:"id"`
	Username string `json:"username" orm:"unique"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Salt string `json:"-"`
	Role string `json:"role"`
	Token string `json:"token"`
	Register time.Time `json:"register_time" orm:"auto_now_add;type(datetime)"`
	LastLogin time.Time `json:"last_login" orm:"auto_now;type(datetime)"`
}
