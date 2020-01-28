package main

import (
	"flag"
	"log"
	"github.com/ZombieMInd/dip_v0.1_go/internal/app/apiserver/apiserver.go"
)

var {
	configPath string
}

func init {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}


func main() {
	config := apiserver.NewConfig()

	s := apiserver.New(config)
	if err = s.Start(); err != nil {
		log.Fatal(err)
	}
}
