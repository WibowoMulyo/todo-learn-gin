package models

type Todo struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}