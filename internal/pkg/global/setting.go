package global

import (
	"simple-demo/pkg/logger"
	"simple-demo/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	S3StoreSetting  *setting.S3StoreSettingS
	Logger          *logger.Logger
)
