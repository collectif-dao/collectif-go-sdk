package main

import (
	"collective-go-sdk/config"
	"collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/logger"
	"collective-go-sdk/worker"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		panic("Unable to read config")
	}

	_ = os.Mkdir("logs", os.ModePerm)
	_ = os.Mkdir("logs/client", os.ModePerm)
	logFile := "./logs/client/" + time.Now().Format("2006-01-02") + ".log"
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFile)
		panic(err)
	}
	defer f.Close()
	logger.InitLogger(f, config.LogLevel)

	ctx := context.Background()
	client, err := fvm.NewLotusClient(ctx, config, keystore.FSKeyStore)
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}

	minerList := config.Miners[config.DefaultNetwork]
	miners, err := worker.PrepareMiners(ctx, minerList, client)
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}

	worker := worker.NewWorker(time.Duration(config.PledgeTimeout)*time.Hour, client, miners, true)

	worker.Start(ctx)
}
