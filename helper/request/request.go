package request

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func Validator(c *gin.Context, rules govalidator.MapData, requestBody interface{}) url.Values {
	opt := govalidator.Options{
		Request: c.Request,
		Data:    requestBody,
		Rules:   rules,
	}

	validator := govalidator.New(opt)
	err := validator.ValidateJSON()
	return err
}
