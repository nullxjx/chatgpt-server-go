package main

import (
	"flag"
	"fmt"

	"github.com/nullxjx/chatgpt-server-go/cmd/chatgpt/app"
	log "github.com/sirupsen/logrus"
)

var configPath = flag.String("config", "./config/config.json", "config file path")

func main() {
	svr, err := app.New(*configPath)
	if err != nil {
		log.Fatalf("create svr err: %v", err)
	}

	if err := svr.HttpEngine.Run(fmt.Sprintf(":%d", svr.Config.HttpPort)); err != nil {
		log.Errorf("failed to run http server: %v", err)
		return
	}
}
