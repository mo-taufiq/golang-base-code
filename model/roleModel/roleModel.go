package roleModel

import (
	"golang-base-code/model"
)

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	model.ModelFieldsDefault
}
