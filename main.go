package main

import (
	"fmt"
	"github.com/dunzoit/go_docker_learn/api"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/create", api.Create)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

