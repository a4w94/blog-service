package main

import (
	"blog_service/global"
	"blog_service/pkg/setting"
	"fmt"
	"log"

	"time"
)

func init() {
	fmt.Println("init")
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

}

func main() {
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
	fmt.Println("global.ServerSetting", global.ServerSetting)

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return nil
	}
	fmt.Println(global.AppSetting)

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return nil
	}
	fmt.Println(global.DatabaseSetting)

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil

}
