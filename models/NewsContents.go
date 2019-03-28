package models

type NewsContents struct {
	Id int `json:"id"`
	Contents string `json:"content"`
	Type string `json:"news_type"`
}
