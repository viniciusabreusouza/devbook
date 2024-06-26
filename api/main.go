package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	fmt.Printf("Rodando API na porta %d", config.Port)

	router := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
