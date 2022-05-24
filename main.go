package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"simple-demo/internal/model"
	global2 "simple-demo/internal/pkg/global"
	"simple-demo/internal/routers"
	"simple-demo/pkg/logger"

	"simple-demo/pkg/setting"

	"simple-demo/pkg/tracer"
	"strings"
	"time"
)

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	// err = setupValidator()
	// if err != nil {
	// 	log.Fatalf("init.setupValidator err: %v", err)
	// }
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:         ":" + global2.ServerSetting.HttpPort,
		Handler:      router,
		ReadTimeout:  global2.ServerSetting.ReadTimeout,
		WriteTimeout: global2.ServerSetting.WriteTimeout,
	}
	global2.Logger.Info(context.Background(), "启动抖音APP服务")
	log.Println("启动抖音APP服务")
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()
	return nil
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global2.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global2.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global2.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global2.JWTSetting)
	if err != nil {
		return err
	}

	global2.AppSetting.DefaultContextTimeout *= time.Second
	global2.JWTSetting.Expire *= time.Second
	global2.ServerSetting.ReadTimeout *= time.Second
	global2.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global2.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global2.ServerSetting.RunMode = runMode
	}

	return nil
}

func setupLogger() error {
	//fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	//global.Logger = logger.NewLogger(&lumberjack.Logger{
	//	Filename:  fileName,
	//	MaxSize:   500,
	//	MaxAge:    10,
	//	LocalTime: true,
	//}, "", log.LstdFlags).WithCaller(2)
	global2.Logger = logger.NewLogger(os.Stdout, "", log.LstdFlags)

	return nil
}

func setupDBEngine() error {
	var err error
	global2.DBEngine, err = model.NewDBEngine(global2.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

// func setupValidator() error {
// 	global.Validator = validator.NewCustomValidator()
// 	global.Validator.Engine()
// 	binding.Validator = global.Validator

// 	return nil
// }

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("simple_demo_tiktok", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global2.Tracer = jaegerTracer
	return nil
}
