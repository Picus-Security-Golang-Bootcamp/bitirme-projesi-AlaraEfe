package main

import (
	"fmt"
	"github.com/AlaraEfe/BasketService/packages/config"
	"log"
)

func main() {
	log.Println("Basket service starting...")

	//Setting environment variables for local development
	cfg, err := config.LoadConfig("./packages/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	fmt.Println(cfg)

}
