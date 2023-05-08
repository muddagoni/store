package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"store/api/v1/order"
	"store/api/v1/product"
	"store/internal/config"
	omodel "store/internal/service/order/model"
	"store/internal/service/product/model"

	"syscall"

	"github.com/gorilla/mux"
)

const (
	configPath = "./config.yaml"
)

func main() {
	config, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	router := mux.NewRouter()
	s := &http.Server{
		Addr:    ":" + fmt.Sprintf("%v", config.Server.Port),
		Handler: router,
	}
	fmt.Printf("http server started on %d ...", config.Server.Port)

	m := make(map[int]*model.Product)
	o := make(map[int]*omodel.OrderRes)

	order.RegisterHandlers(router, m, o)
	product.RegisterHandlers(router, m)

	//gracefull shutdown
	connClose := make(chan struct{})
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sig

		err = s.Shutdown(context.Background())
		if err != nil {
			log.Panic(err.Error())
		}
		close(connClose)
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server shut down error %v", err)
	}
	<-connClose
	fmt.Println("server stopped...")
}
