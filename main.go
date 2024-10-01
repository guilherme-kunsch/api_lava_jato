package main

import (
	"fmt"
	"lavajato/src/config"
	"lavajato/src/router"
	"log"
	"net/http"
)

func main() {
	config.ToLoad()

	r := router.ToGenerate()

	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
