package db

import (
	"go-web/global/model"
	"gorm.io/gorm"
	"reflect"
)

var (
	GVA_DB *CustomDB
)

type CustomDB struct {
	*gorm.DB
}

func (db *CustomDB) Page(tx *gorm.DB, pageQuery model.PageQueryModel, dbModeType interface{}) *model.PageModel {

	//1.默认参数
	if pageQuery.Size < 1 {
		pageQuery.Size = 10
	}
	if pageQuery.Page < 1 {
		pageQuery.Page = 1
	}

	typeOfT := reflect.TypeOf(dbModeType)
	dbModel := reflect.New(typeOfT).Interface()
	sliceOfT := reflect.SliceOf(typeOfT)
	list := reflect.MakeSlice(sliceOfT, 0, 0).Interface()
	var total int64
	tx.Model(&dbModel).Count(&total)

	pageModel := &model.PageModel{
		List:  &list,
		Total: &total,
		Page:  pageQuery.Page,
		Size:  pageQuery.Size,
	}
	offset := pageQuery.Size * (pageQuery.Page - 1)
	tx.Limit(pageModel.Size).Offset(offset).Find(&list)

	pageModel.List = &list
	return pageModel
}
