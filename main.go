package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/20gu00/aBais/common/config"
	initDo "github.com/20gu00/aBais/common/init-do"
	"github.com/20gu00/aBais/dao/db"

	"go.uber.org/zap"
)

func main() {
	// 初始化
	r := initDo.InitDo()

	srv := &http.Server{
		Addr:           config.Config.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(config.Config.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Config.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << config.Config.MaxHeader, // 1的xx次方
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()
	fmt.Println("[Info] server port ", strings.Split(config.Config.Addr, ":")[1])

	quit := make(chan os.Signal, 2)
	// interrupt中断信号 syscall.SIGTERM, syscall.SIGINT
	signal.Notify(quit, os.Interrupt)
	// 空则阻塞,监听第一次中断信号,用于优雅关闭
	<-quit

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.GraceTime)*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			zap.L().Fatal("Gin Server关闭异常:", zap.Error(err))
		}
		zap.L().Info("Gin Server成功退出")

		if err := db.DBClose(); err != nil {
			zap.L().Fatal("DB关闭异常:", zap.Error(err))
		}
	}()

	go func() {
		timerC := time.NewTimer(time.Duration(config.Config.GraceTime) * time.Second).C
		<-timerC
		fmt.Println("程序正常退出完毕")
		os.Exit(0)
	}()

	// 第二次中断信号,直接退出
	<-quit
	os.Exit(1)
}
