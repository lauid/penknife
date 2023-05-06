package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func SignalListen() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
