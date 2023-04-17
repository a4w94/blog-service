package main

import (
	"blog_service/global"
	"blog_service/internal/model"
	"blog_service/pkg/logger"
	"blog_service/pkg/setting"
	"log"
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
		log.Fatalf("init.setupLooger err: %v", err)
	}

}

func main() {

	global.Logger.Infof("%s: go-programing-tour-book/%s", "eddycjy", "blog-service")

	// router := routers.NewRouter()
	// s := &http.Server{
	// 	Addr:           ":" + global.ServerSetting.HttpPort,
	// 	Handler:        router,
	// 	ReadTimeout:    global.ServerSetting.ReadTimeout,
	// 	WriteTimeout:   global.ServerSetting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

}

func setupSetting() error {
	setting, err := setting.NetSetting()
	if err != nil {
		return nil
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return nil
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return nil
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return nil
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

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
	global.Logger = logger.NewLogger(writer, "", log.LstdFlags).WithCaller(2)

	return nil

}
