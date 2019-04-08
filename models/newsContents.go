package models
// 新闻详情表
type NewsContents struct {
	Id int `json:"id"`
	Contents string `json:"content"`
	Type string `json:"news_type"`
}
