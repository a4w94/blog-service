package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// 參數綁定&導入參數驗證
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors

	//c.ShouldBind()方法，該方法從請求的form或querystring中綁定值
	//c.ShouldBind 則會根據 Content-Type Header，自動選擇對應的綁定器，可以用於綁定 JSON、XML、ProtoBuf、Form、Query 等多種格式的請求內容
	err := c.ShouldBind(v)

	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		//對validError進行二次封裝
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}

		//對錯誤訊息進行翻譯
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
