package book

import (
	"go-web/global/model"
)

type Book struct {
	model.Model
	Name        string  `json:"name"      gorm:"type:varchar(50);comment:书名" `
	Author      string  `json:"author"    gorm:"type:varchar(20);comment:作者" `
	ChiefEditor string  `json:"chiefEditor"  gorm:"type:varchar(20);comment:主编" `
	Price       float64 `json:"price"  gorm:"type:decimal(20,2);comment:单价" `
}
