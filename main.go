package main

import (
	"blog_service/global"
	"blog_service/internal/model"
	"blog_service/internal/routers"
	"blog_service/pkg/email"
	"blog_service/pkg/logger"
	"blog_service/pkg/setting"
	"blog_service/pkg/tracer"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/natefinch/lumberjack/v3"
)

var (
	port    string
	runMode string
	config  string
)

var (
	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
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
	fmt.Printf("isVersion: %v\n", isVersion)
	if isVersion {
		fmt.Printf("buildTime: %s\n", buildTime)
		fmt.Printf("buildVersion: %s\n", buildVersion)
		fmt.Printf("gitCommitID: %s\n", gitCommitID)
		return
	}

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

	setting, err := setting.NetSetting(strings.Split(config, ",")...)
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

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}

	err = setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}

	if port != "" {
		global.ServerSetting.HttpPort = port
	}

	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil

}

// DB設定
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}

// Log檔存放位址設定
func setupLogger() error {

	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	writer, err := lumberjack.NewRoller(fileName, int64(global.AppSetting.LogMaxSize), &lumberjack.Options{
		MaxAge:    10,
		LocalTime: true,
	})
	if err != nil {
		log.Fatal("lumberjack failed")
	}
	global.Logger = logger.NewLogger(writer, "", log.LstdFlags).WithCaller(1)
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

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "啟動通訊埠")
	flag.StringVar(&runMode, "mode", "", "啟動模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的設定檔路徑")
	flag.BoolVar(&isVersion, "version", false, "編譯資訊")
	flag.Parse()
	return nil
}
