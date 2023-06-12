package setting

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// config相關設定
func NetSetting(configs ...string) (*Setting, error) {
	fmt.Println("NewSetting")
	vp := viper.New()

	vp.SetConfigName("config")

	for _, config := range configs {
		//判斷設定檔是否存在
		if config != "" {
			vp.AddConfigPath(config)
		}
	}

	//* 使用此方法是因為使用go build沒有打包config檔，而使用路徑去尋找
	//* 其他解決方法可使用go-bindata或go env進行系統環境變數設定
	//* 一般來說系統環境變數讀取：命令列參數>系統環境變數>檔案設定
	//設定config路徑
	_, filename, _, ok := runtime.Caller(1) //回應main.go檔案位置
	if !ok {
		log.Fatal("setting.NewSetting error,runtime.Caller error")
	}
	dir := filepath.Dir(filename)               //回應main.go暫存根目錄
	configPath := filepath.Join(dir, "configs") //設定configs目錄路徑
	vp.AddConfigPath(configPath)

	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()

	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

// 熱更新
func (s *Setting) WatchSettingChange() {
	go func() {
		//對檔案設定進行監聽
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			log.Println("Config changed:", in.Name, in.Op)
			_ = s.ReloadAllSection()
		})
	}()
}
