package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
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

	quit := make(chan os.Signal)
	// syscall.SIGTERM, syscall.SIGINT
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Gin Server关闭异常:", zap.Error(err))
	}
	zap.L().Info("Gin Server成功退出")

	if err := db.DBClose(); err != nil {
		zap.L().Fatal("DB关闭异常:", zap.Error(err))
	}
}
