package v1

import (
	"blog_service/global"
	"blog_service/internal/service"
	"blog_service/pkg/app"
	"blog_service/pkg/convert"
	"blog_service/pkg/errorcode"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Tag struct {
	Name    string      //標籤名稱
	Content interface{} //暫存欄位
}

func NewTag() Tag {
	return Tag{}
}

func (r Tag) Get(c *gin.Context) {

}

// @Summary 取得多個標籤
// @Produce json
// @Param name query string false "標籤名稱" maxlength(100)
// @Param state query int false "狀態" Enums(0,1) default(1)
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errorcode.Error "請求錯誤"
// @Failure 500 {object} errorcode.Error "內部錯誤"
// @Router /api/v1/tags [get]
func (r Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)

	//參數校正碼綁定
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errorcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	//取得標籤總數
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag error: %v", err)
		response.ToErrorResponse(errorcode.ErrorCountTagFail)
		return
	}

	//取得標籤列表
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList error: %v", err)
		response.ToErrorResponse(errorcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)

}

// @Summary 新增標籤
// @Produce json
// @Param name query string true "標籤名稱" minlenght(3) maxlength(100)
// @Param state query int false "狀態" Enums(0,1) default(1)
// @Param created_by query  string true "建立者" minlenght(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errorcode.Error "請求錯誤"
// @Failure 500 {object} errorcode.Error "內部錯誤"
// @Router /api/v1/tags [post]
func (r Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		errRsp := errorcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag error: %v", err)
		response.ToErrorResponse(errorcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})

}

// @Summary 更新標籤
// @Produce json
// @Param id path int true "標籤ID"
// @Param name body string false "標籤名稱" minlenght(3) maxlength(100)
// @Param state body int false "狀態" Enums(0,1) default(1)
// @Param modified_by body  string true "修改者"minlenght(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errorcode.Error "請求錯誤"
// @Failure 500 {object} errorcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [put]
func (r Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo.MustUInt32(convert.StrTo(c.Param("id"))),
	}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	fmt.Println(param.ID)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		errRsp := errorcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag error: %v", err)
		response.ToErrorResponse(errorcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 刪除標籤
// @Produce json
// @Param id path int true "標籤ID"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errorcode.Error "請求錯誤"
// @Failure 500 {object} errorcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [delete]
func (r Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo.MustUInt32(convert.StrTo(c.Param("id"))),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		errRsp := errorcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag error: %v", err)
		response.ToErrorResponse(errorcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
}
