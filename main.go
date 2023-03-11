package main

import (
	"github.com/voyager-go/GoWeb/cmd"
	"github.com/voyager-go/GoWeb/pkg/logging"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		logging.Log.Error(err)
	}
	// 等待程序退出信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		// 接收到信号后退出程序
		os.Exit(0)
	}()
}
