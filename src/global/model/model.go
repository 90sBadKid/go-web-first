package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type PageModel struct {
	List  *interface{} `json:"list"`
	Total *int64       `json:"total"`
	Page  int          `json:"page"`
	Size  int          `json:"size"`
}

type PageQueryModel struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
