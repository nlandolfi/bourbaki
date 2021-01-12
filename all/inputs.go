package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

type Entry struct {
	Title string
	Name  string
}

func main() {
	tmpl := template.Must(
		template.ParseFiles("inputs.tmpl"),
	)

	f, err := os.Open("sheets.csv")
	if err != nil {
		log.Fatal("os.Open: %v", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	var entries []Entry
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

		log.Print(record[0])
		entries = append(entries,
			Entry{
				Title: strings.Title(strings.Join(strings.Split(record[0], "_"), " ")),
				Name:  record[0],
			},
		)
	}

	tmpl.Execute(os.Stdout, entries)
}
