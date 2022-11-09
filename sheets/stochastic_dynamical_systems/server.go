package main

import (
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/", http.FileServer(http.Dir("./")))
	log.Fatal(http.ListenAndServe(":8088", m))
}
