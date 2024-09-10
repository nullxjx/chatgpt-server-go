package main

import (
	"flag"
	"fmt"

	"github.com/nullxjx/chatgpt-server-go/cmd/chatgpt/app"
	log "github.com/sirupsen/logrus"
)

func startHttpServer(app *app.App, done chan<- error) {
	if err := app.HttpEngine.Run(fmt.Sprintf(":%d", app.Config.HttpPort)); err != nil {
		log.Errorf("failed to run http server: %v", err)
		done <- err
		return
	}
	done <- nil
}

var configPath = flag.String("config", "./config/config.json", "config file path")

func main() {
	svr, err := app.New(*configPath)
	if err != nil {
		log.Fatalf("create svr err: %v", err)
	}
	done := make(chan error, 1)
	go startHttpServer(svr, done)
	if err := <-done; err != nil {
		log.Fatalf("init svr err: %v", err)
		return
	}
}
