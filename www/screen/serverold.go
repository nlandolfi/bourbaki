package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	staticDir = flag.String("static-dir", "/static", "directory of static files")
)

func main() {
	flag.Parse()
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(*staticDir))))
}
