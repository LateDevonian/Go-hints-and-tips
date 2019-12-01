package main

import (
	"log/syslog"
)

func main() {
	//set syslog client
	// the faclity to log to is log_local3
	//the prefix you want the message to start with narwhal
	//shoutl be the service name
	logger, err := syslog.New(syslog.LOG_LOCAL3, "narwhal")
	if err != nil {
		panic("cannot attach to syslog")
	}
	defer logger.Close()

	logger.Debug("Debug message")
	logger.Notice("Notice message")
	logger.Warning("warning message")
	logger.Alert("Alert message")
}
