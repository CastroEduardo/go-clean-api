package validation

import (
	"github.com/go-playground/validator/v10"
)

func IranianMobileNumberValidator(fld validator.FieldLevel) bool {

	// value, ok := fld.Field().Interface().(string)
	// if !ok {
	// 	return false
	// }

	return true // common.IranianMobileNumberValidate(value)
}
