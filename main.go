package main

import (
	"flag"

	"github.com/rayguo17/go-warmup/config"
	"github.com/rayguo17/go-warmup/logger"
	"go.uber.org/zap"
)

var configFile string

func init() {
	config.Init()
}

func main() {

	if !flag.Parsed() {
		flag.Parse()
	}
	config.ReadConfig()
	url := "test"
	logger.Logger.Info("initialize done!",
		zap.String("url", url),
	)
	//create a key-value store support multi-thread access.
}

//listen to socket, access connection, and then read command to change
//
func MainRountine() {

}
