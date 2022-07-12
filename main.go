package main

import (
	"flag"
	"fmt"
	"github.com/Gary-Gs/go-clean-arch/app"
	"log"
	"os"
	"strings"

	_ "github.com/Gary-Gs/go-clean-arch/resources/webapps/swagger"
)

// @title Golang Service Blueprint
// @version 1.0
// @description boilerplate code for backend service in golang
// @BasePath /api/v1/
func main() {
	env := os.Getenv("DEPLOYMENT_ENV")
	var configFile *string
	switch strings.ToUpper(env) {
	case "BETA":
		configFile = flag.String("config", "config-beta.yaml", "relative/absolute config file path")
	default:
		configFile = flag.String("config", "config.yaml", "relative/absolute config file path")
	}
	fmt.Println("Starting application with config file: ", *configFile)
	flag.Parse()

	a, err := app.GetNewInstance(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(a.InitApiEcho().Start(a.Configs.AppConfig.ApiPort))
}
