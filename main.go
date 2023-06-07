package main

import (
	"blog_service/global"
	"blog_service/internal/model"
	"blog_service/internal/routers"
	"blog_service/pkg/email"
	"blog_service/pkg/logger"
	"blog_service/pkg/setting"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/natefinch/lumberjack/v3"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

}

// @title 部落格系統
// @version 1.0
// @description 部落格建立
// @termsOfService https://github.com/go-programing-tour-book
func main() {

	//global.Logger.Infof("%s: go-programing-tour-book/%s", "eddycjy", "blog-service")

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Url: http://127.0.0.1:%s", global.ServerSetting.HttpPort)
	s.ListenAndServe()

}

func setupSetting() error {
	setting, err := setting.NetSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout *= time.Second

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	fmt.Println("啟用email 測試")
	//emailTest()

	return nil

}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {

	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	writer, err := lumberjack.NewRoller(fileName, 600, &lumberjack.Options{
		MaxAge:    10,
		LocalTime: true,
	})
	if err != nil {
		log.Fatal("lumberjack failed")
	}
	global.Logger = logger.NewLogger(writer, "", log.LstdFlags).WithCaller(4)

	return nil

}

func emailTest() {
	defailMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	},
	)

	err := defailMailer.SendEmail(
		global.EmailSetting.To,
		fmt.Sprintf("例外拋出，發生時間：%d", time.Now().Unix()),
		fmt.Sprintf("錯誤訊息：%v"),
	)

	if err != nil {
		panic(err)
	}
}
