package middleware

import "github.com/gin-gonic/gin"

//處理程序內上下文設定內部資訊
// Ex: 可在處理器函數中使用c.GetString("app_name")獲得值..
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog-service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
