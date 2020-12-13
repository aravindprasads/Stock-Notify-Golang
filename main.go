package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/new", AddHandler)
	http.HandleFunc("/delete", DeleteHandler)
	fmt.Printf("Server started at %d", 8080)
	go timer_main()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
