package page

import (
	"gorm.io/gorm"
	"reflect"
)

type PageModel struct {
	List  *interface{} `json:"list"`
	Total *int64       `json:"total"`
}

type PageQueryModel struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func Paginate(db *gorm.DB, pageQuery *PageQueryModel, dbModeType interface{}) *PageModel {

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
	db.Model(&dbModel).Count(&total)

	pageModel := &PageModel{
		List:  &list,
		Total: &total,
	}

	offset := pageQuery.Size * (pageQuery.Page - 1)
	db.Limit(pageQuery.Page).Offset(offset).Find(&list)

	pageModel.List = &list
	return pageModel
}
