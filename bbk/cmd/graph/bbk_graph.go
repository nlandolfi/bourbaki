package main

import (
	"bbk"
	"flag"
	"log"
	"os"
	"text/template"
)

var (
	inputFileFlag    = flag.String("csv", "graph.csv", "csv file")
	templateFileFlag = flag.String("tmpl", "graph.tmpl", "which template")
)

func main() {
	flag.Parse()

	tmpl := template.Must(
		template.ParseFiles(*templateFileFlag),
	)

	f, err := os.Open(*inputFileFlag)
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}
	defer f.Close()

	entries, err := bbk.ParseGraph(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.Execute(os.Stdout, entries); err != nil {
		log.Fatal(err)
	}
}
