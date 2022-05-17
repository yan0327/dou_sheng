package global

import (
	"simple-demo/pkg/logger"
	"simple-demo/pkg/setting"

	"golang.org/x/sync/singleflight"
)

var (
	ServerSetting           *setting.ServerSettingS
	AppSetting              *setting.AppSettingS
	JWTSetting              *setting.JWTSettingS
	DatabaseSetting         *setting.DatabaseSettingS
	Logger                  *logger.Logger
	GVA_Concurrency_Control = &singleflight.Group{}
)
