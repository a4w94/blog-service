package v1

import (
	"blog_service/global"
	"blog_service/internal/service"
	"blog_service/pkg/app"
	"blog_service/pkg/errorcode"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}

	c.ShouldBind(&param)
	fmt.Println("param:", param)
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		global.Logger.Infof("params %v", param)
		errRsp := errorcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	//判斷是否存在
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.ChckcAuth err: %v", err)
		response.ToErrorResponse(errorcode.UnauthorizedAuthNotExist)
		return
	}

	//產生token
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errorcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})

}
