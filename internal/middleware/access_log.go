package middleware

import (
	"blog_service/global"
	"blog_service/pkg/logger"
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 捕獲響應內容進行處理
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

// 寫入紀錄檔
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix() //請求開始時間
		c.Next()
		endTime := time.Now().Unix() //請求結束時間

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(), //請求參數
			"response": bodyWriter.body.String(),    //請求結果回應主體
		}

		//!log 使用WithFields無法生成log
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d,begin_time: %d, end_time: %d",
			c.Request.Method, //呼叫方法
			bodyWriter.Status(),
			beginTime,
			endTime,
		)

	}
}
