package jsonHelper

import (
	"encoding/json"
	"fmt"

	gologger "github.com/mo-taufiq/go-logger"
)

func ConvertStructToTidyJSON(s interface{}) string {
	byt, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		gologger.Error(fmt.Sprintf("error convert struct to tidy JSON: %s", err.Error()))
	}
	return string(byt)
}
