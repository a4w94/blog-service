package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// 逾時控制中介軟體
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		//設定context的逾時時間
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()

	}
}
