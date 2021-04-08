package translator

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin/binding"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/sirupsen/logrus"
)

var v *validator.Validate
var trans ut.Translator

func Init() {
	// 中文翻译
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	var found bool
	trans, found = uni.GetTranslator("zh")
	if !found {
		logrus.Error("验证器找不到中文语言包")
		return
	}

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 验证器注册翻译器
		zh_translations.RegisterDefaultTranslations(v, trans)
	}
}

func Trans(err error) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	var errList []string
	for _, e := range errs {
		errList = append(errList, e.Translate(trans))
	}
	return errors.New(strings.Join(errList, "|"))
}
