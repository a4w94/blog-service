package email

import (
	"blog_service/global"
	"blog_service/pkg/setting"
	"fmt"

	"testing"
)

func TestEmail_SendEmail(t *testing.T) {
	fmt.Println("test")
	setting, err := setting.NetSetting()
	if err != nil {
		fmt.Println(err)
	}
	setting.ReadSection("Email", &global.EmailSetting)
	// e := NewEmail(&SMTPInfo{
	// 	Host:     global.EmailSetting.Host,
	// 	Port:     global.EmailSetting.Port,
	// 	IsSSL:    global.EmailSetting.IsSSL,
	// 	UserName: global.EmailSetting.UserName,
	// 	Password: global.EmailSetting.Password,
	// 	From:     global.EmailSetting.From,
	// })
	// fmt.Println("e", e)

}
