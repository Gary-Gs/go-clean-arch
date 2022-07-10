package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	env := os.Getenv("DEPLOYMENT_ENV")
	var configFile *string
	switch strings.ToUpper(env) {
	case "BETA":
		configFile = flag.String("config", "config/config-beta.yaml", "relative/absolute config file path")
	default:
		configFile = flag.String("config", "config/config.yaml", "relative/absolute config file path")
	}
	fmt.Println("Starting application with config file: ", *configFile)
	flag.Parse()

	a, err := GetNewInstance(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(a.InitApiEcho().Start(a.Configs.AppConfig.ApiPort))
}
