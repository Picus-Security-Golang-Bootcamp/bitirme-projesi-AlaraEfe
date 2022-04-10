package main

import (
	"fmt"
	"github.com/AlaraEfe/BasketService/packages/config"
	"github.com/AlaraEfe/BasketService/packages/db"
	"github.com/AlaraEfe/BasketService/packages/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Basket service starting...")

	//Setting environment variables for local development
	cfg, err := config.LoadConfig("./packages/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	fmt.Println(cfg)

	// Seting Global Logger then close it
	logger.NewLogger(cfg)
	defer logger.Close()

	db := db.ConnectDB(cfg)

	router := gin.Default()

	server := &http.server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

}
