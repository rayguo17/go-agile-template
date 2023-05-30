package config

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/rayguo17/go-warmup/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var configFile string

type conf struct {
	LogConf zap.Config `yaml:"log"`
}

var Config conf

func ReadConfig() {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	//pp.Println(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	// pp.Println(Config)
	// pp.Println(Config.LogConf.EncoderConfig)
	logger.Logger, err = Config.LogConf.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Logger.Sync()
}

func cmdlineConfig() {
	flag.StringVar(&configFile, "config", "./config.yaml", "indicate the path of config.yaml")
}

func Init() {
	cmdlineConfig()
	// fmt.Println(configFile)
}
