package main

import (
	"log"

	"github.com/spf13/viper"
)

type rpcServiceConfig struct {
	Host    string `json:host`
	Timeout int    `json:timeout`
}

type rpcServices struct {
	services map[string]rpcServiceConfig `json:rpcServices`
}

var rpcServicesConfig rpcServices

func readConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	log.Printf("%v\n", viper.Get("rpcServices"))

	err = viper.UnmarshalKey("rpcServices", &rpcServicesConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func main() {
	readConfig()

	log.Printf("Environemnt: %v\n", viper.Get("environment"))
	log.Printf("RPC services: %v\n", rpcServicesConfig)
}
