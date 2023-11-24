package main

import (
	"log"
	"net/http"

	"github.com/Isomiddinov7/exam/config"
	"github.com/Isomiddinov7/exam/controller"
	"github.com/Isomiddinov7/exam/storage/postgres"
)

func main() {

	var cfg = config.Load()

	pgStorage, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	handler := controller.NewController(&cfg, pgStorage)

	http.HandleFunc("/category", handler.Category)
	http.HandleFunc("/product", handler.Product)
	http.HandleFunc("/branch", handler.Branch)
	http.HandleFunc("/client", handler.Client)
	http.HandleFunc("/order", handler.Order)

	log.Println("Listening:", cfg.ServiceHost+cfg.ServiceHTTPPort, "...")
	if err := http.ListenAndServe(cfg.ServiceHost+cfg.ServiceHTTPPort, nil); err != nil {
		panic("Listent and service panic:" + err.Error())
	}
}
