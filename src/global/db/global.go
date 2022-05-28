package db

import (
	"gorm.io/gorm"
)

var (
	GVA_DB *CustomDB
)

type CustomDB struct {
	*gorm.DB
}
