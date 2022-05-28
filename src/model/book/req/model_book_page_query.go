package req

import (
	"go-web/global/page"
)

type ModelBookPageQuery struct {
	*page.PageQueryModel

	Name string `json:"name"`
}
