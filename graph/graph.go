package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	"github.com/alecthomas/template"
)

type Entry struct {
	Name  string
	Needs []string
}

var templateFileFlag = flag.String("tmpl", "graph.tmpl", "which template")

func main() {
	flag.Parse()

	tmpl := template.Must(
		template.ParseFiles(*templateFileFlag),
	)

	f, err := os.Open("graph.csv")
	if err != nil {
		log.Fatal("os.Open: %v", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	entries := make(map[string]*Entry)
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

		name := record[0]
		entry, ok := entries[name]
		if !ok {
			entry = &Entry{
				Name: name,
			}
			entries[name] = entry
		}
		need := record[1]
		entry.Needs = append(entry.Needs, need)
	}

	tmpl.Execute(os.Stdout, entries)
}
