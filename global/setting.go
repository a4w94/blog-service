package global

import (
	"blog_service/pkg/logger"
	"blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)
