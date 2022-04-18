package main

import (
	"fmt"
	"github.com/AlaraEfe/BasketService/internal/basket"
	"github.com/AlaraEfe/BasketService/internal/category"
	"github.com/AlaraEfe/BasketService/internal/product"
	"github.com/AlaraEfe/BasketService/internal/user"
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

	//fmt.Println(cfg)

	// Seting Global Logger then close it
	logger.NewLogger(cfg)
	defer logger.Close()

	db := db.ConnectDB(cfg)

	router := gin.Default()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

	rootRouter := router.Group(cfg.ServerConfig.RoutePrefix)
	productRouter := rootRouter.Group("/product")
	categoryRouter := rootRouter.Group("/category")
	userRouter := rootRouter.Group("/user")
	basketRouter := rootRouter.Group("/basket")

	//Product Repository
	productRepo := product.NewProductRepository(db)
	productRepo.Migration()
	product.NewProductHandler(productRouter, productRepo, cfg)

	//Category Repository
	categoryRepo := category.NewCategoryRepository(db)
	categoryRepo.Migration()
	category.NewCategoryHandler(categoryRouter, categoryRepo, cfg)

	//User Repository
	userRepo := user.NewUserRepository(db)
	userRepo.Migration()
	user.NewUserHandler(userRouter, userRepo, cfg)

	//BasketItem Repository
	basketItemRepo := basket.NewBasketItemRepository(db)
	basketItemRepo.Migration()
	basket.NewBasketItemHandler(basketRouter, basketItemRepo, cfg)

	router.Run()

	go func() {

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}()

	log.Println("Book store service started")

}
