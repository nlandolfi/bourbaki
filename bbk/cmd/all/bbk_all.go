package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"bbk"
)

var (
	sheetsDir = flag.String("sheets", "../", "where to find other sheets")
)

func main() {
	flag.Parse()

	results, err := bbk.ParseAll(*sheetsDir)
	if err != nil {
		log.Fatal(err)
	}

	rs := make(map[string]*bbk.ParseResult, len(results))
	for _, p := range results {
		rs[p.Config.Name] = p
	}

	// get the template
	order := mustGetAllSheetsOrder("sheets.csv")
	orderedResults := make([]*bbk.ParseResult, len(order))
	for i, n := range order {
		r, ok := rs[n]
		if !ok {
			panic(fmt.Sprintf("unknown sheet: %v", n))
		}
		orderedResults[i] = r
	}

	inputsTemplate := template.Must(
		template.New("inputs.tmpl").Funcs(
			template.FuncMap{
				"title": bbk.Title,
				"join": func(ss []string) string {
					return strings.Join(ss, ", ")
				},
			},
		).ParseFiles("inputs.tmpl"))

	if err := inputsTemplate.Execute(os.Stdout, orderedResults); err != nil {
		log.Fatal(err)
	}
}

func mustGetAllSheetsOrder(filename string) []string {
	// get the order
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("os.Open: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	var order []string
	// var entries []Entry
	var header = true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if header {
			header = false
			continue
		}

		order = append(order, record[0])
	}
	return order
}
