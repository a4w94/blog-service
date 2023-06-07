package middleware

import (
	"blog_service/global"
	"blog_service/pkg/app"
	"blog_service/pkg/email"
	"blog_service/pkg/errorcode"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 自訂recovery
func Recovery() gin.HandlerFunc {
	defailMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	},
	)
	
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover error: %v"
				global.Logger.WithCallerFrames().Errorf(s, err)

				err := defailMailer.SendEmail(
					global.EmailSetting.To,
					fmt.Sprintf("例外拋出，發生時間：%d", time.Now().Unix()),
					fmt.Sprintf("錯誤訊息：%v", err),
				)

				if err != nil {
					global.Logger.Panicf("mail.SendEmail err: %v", err)

					app.NewResponse(ctx).ToErrorResponse(errorcode.ServerError)
					ctx.Abort()
				}
			}
		}()
		ctx.Next()
	}
}
