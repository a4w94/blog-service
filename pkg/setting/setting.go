package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NetSetting() (*Setting, error) {
	fmt.Println("NewSetting")
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return &Setting{vp}, err
}
