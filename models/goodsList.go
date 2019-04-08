package models

import "Vue-cms-server/lib"

type GoodsList struct {
	Id int `json:"id"`
	Title string `json:"title"`
	AddTime lib.Time `json:"add_time"`
	Content string `json:"content"`
	Click int `json:"click"`
	ImgUrl string `json:"img_url"`
	SellPrice int `json:"sell_price"`
	MarketPrice int `json:"market_price"`
	StockQuantity int `json:"stock_quantity"`
	GoodsSwipe []*GoodsSwipe `json:"-" orm:"reverse(many)"`
}
