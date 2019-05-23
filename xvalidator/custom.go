package xvalidator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/globalsign/mgo/bson"
	"gopkg.in/go-playground/validator.v9"
)

func RegisterCustomValidator() {
	v := binding.Validator.Engine().(*validator.Validate)
	v.RegisterValidation("is-objectid", ValidateObjectId)
}

func ValidateObjectId(fl validator.FieldLevel) bool {
	if idStr, ok := fl.Field().Interface().(string); ok {
		if idStr == "" || bson.IsObjectIdHex(idStr) {
			return true
		}
	}
	return false
}
