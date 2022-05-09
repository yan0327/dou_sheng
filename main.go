package main

import (
	"log"
	"net/http"
	"simple-demo/internal/routers"
	"time"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("启动抖音APP服务")
	s.ListenAndServe()
}
