package req

import "go-web/global/model"

type ModelBookPageQuery struct {
	model.PageQueryModel

	Name string `json:"name"`
}
