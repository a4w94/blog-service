package setting

import (
	"time"
)

type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettings struct {
	DefaultPageSize       int
	MaxPageSize           int
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	LogMaxSize            int
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
	DefaultContextTimeout time.Duration
}

type DatabaseSettings struct {
	DBType       string
	UserName     string
	PassWord     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
type JWTSettings struct {
	Secret string
	//簽發者
	Issuer string
	//簽證留存時間
	Expire time.Duration
}

type EmailSettings struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

var sections = make(map[string]interface{})

// 讀取配置文件並解析
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	//增加讀取section的存儲紀錄，以便在重新載入設定的方法中進行二次處理
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

// 重新讀取設定
func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
