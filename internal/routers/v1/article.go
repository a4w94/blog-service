package v1

import (
	"blog_service/global"
	"blog_service/internal/service"
	"blog_service/pkg/app"
	"blog_service/pkg/errorcode"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errorcode.ServerError)

}

func (a Article) List(c *gin.Context) {

}

// @Summary 新增文章
// @Produce  json
// @Accept json
// @Param createArticleRequest body service.CreateArticleRequest true "文章資料"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errorcode.Error "請求錯誤"
// @Failure 500 {object} errorcode.Error "內部錯誤"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	fmt.Println("CreateArticle")
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)

	//參數校正綁定
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err:%v", errs)
		errRsp := errorcode.InvalidParams.WithDetails(errs.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	fmt.Println("param", param)

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err:%v", err)
		response.ToErrorResponse(errorcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})

}
func (a Article) Update(c *gin.Context) {

}
func (a Article) Delete(c *gin.Context) {

}
