package app

import (
	"blog_service/global"
	"blog_service/pkg/convert"

	"github.com/gin-gonic/gin"
)

//分頁處理

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()

	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

// 计算分页查询时需要跳过的记录数量
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
