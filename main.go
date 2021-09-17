package main

import (
	"fmt"
	"net/http"
	"pro/app/database"
	"pro/app/router"
	"pro/config"
	"time"
)

func main() {
	config.Run()
	if err := database.Run(); err != nil {
		fmt.Println("数据库链接失败:", err)
		return
	}
	if err := database.RedisInit(); err != nil {
		fmt.Println("Reids链接失败:", err)
		return
	}
	r := router.RouteInit()
	r.Static("/static", "./static")
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.HTTPPort),
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 10 << 10,
	}
	s.ListenAndServe()
}
