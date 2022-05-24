package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	gormopentracing "gorm.io/plugin/opentracing"
	"simple-demo/internal/pkg/global"
	"simple-demo/pkg/setting"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	logMode := logger.Warn
	if global.ServerSetting.RunMode == "debug" {
		logMode = logger.Info
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(s,
			databaseSetting.UserName,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		),
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	if logMode >= logger.Info {
		db.Use(gormopentracing.New(
			gormopentracing.WithLogResult(true),
			gormopentracing.WithSqlParameters(true),
		))
	} else {
		db.Use(gormopentracing.New())
	}
	return db, nil
}
