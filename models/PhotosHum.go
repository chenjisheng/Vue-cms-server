package models

type PhotosHum struct {
	Id int `json:"id"`
	HumUrl string `json:"src"`
	PhotoList *PhotosList `orm:"rel(fk)" json:"-"`
}
