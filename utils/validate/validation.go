package validate

import (
	"fmt"
	"reflect"

	"github.com/panda8z/eeyoo/utils/errors"
	"github.com/go-playground/locales/zh_Hans_CN"
	trans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh "github.com/go-playground/validator/v10/translations/zh"
)

func Validate(data interface{}) (string, int) {
	v := validator.New()
	u := trans.New(zh_Hans_CN.New())
	t, _ := u.GetTranslator(zh_Hans_CN.New().Locale())
	err := zh.RegisterDefaultTranslations(v, t)
	if err != nil {
		fmt.Println(err.Error())
	}

	// handel sturct tag label
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = v.Struct(data)
	if err != nil {
		for _, val := range err.(validator.ValidationErrors) {
			return val.Translate(t), errors.ERROR
		}
	}
	return "", errors.SUCCESS
}
