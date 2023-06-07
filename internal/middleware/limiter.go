package middleware

import (
	"blog_service/pkg/app"
	"blog_service/pkg/errorcode"
	"blog_service/pkg/limiter"

	"github.com/gin-gonic/gin"
)

// 限流器中介軟體
func RateLimiter(l limiter.LimitIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1) //回傳值為刪除的權杖數
			//沒有可用權杖=>超出配額
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errorcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
