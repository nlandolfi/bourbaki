package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/template"
)

type Entry struct {
	Name    string
	DirName string
	Needs   []string
}

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
				Name:    name,
				DirName: dirname(name),
			}
			entries[name] = entry
		}
		need := record[1]
		entry.Needs = append(entry.Needs, need)
	}

	tmpl.Execute(os.Stdout, entries)
}

func dirname(s string) string {
	pieces := strings.Split(s, " ")
	for i, p := range pieces {
		pieces[i] = strings.ToLower(p)
	}
	return strings.Join(pieces, "_")
}
