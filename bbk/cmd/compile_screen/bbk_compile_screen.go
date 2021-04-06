package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"text/template"

	"bbk"
)

var (
	sheetsDir = flag.String("sheets-dir", "../sheets", "the sheets directory")
)

func main() {
	flag.Parse()

	results, err := bbk.ParseAll(*sheetsDir)
	if err != nil {
		log.Fatal(err)
	}

	names := make([]string, 0, len(results))
	for n := range results {
		names = append(names, n)
	}
	sort.Strings(names)

	log.Printf("%d sheets", len(results))

	sheetTemplate := template.Must(
		template.New("sheet.tmpl").Funcs(
			template.FuncMap{
				"title": bbk.Title,
			},
		).ParseFiles(path.Join("./static", "sheet.tmpl")))

	for _, name := range names {
		p := results[name]
		out, err := os.Create(fmt.Sprintf("./static/sheets/%s.html", name))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		if err := sheetTemplate.Execute(out, p); err != nil {
			log.Fatal(err)
		}
		out.Close()

		_, err = bbk.CopyFile(
			path.Join(*sheetsDir, fmt.Sprintf("%s/%s.pdf", name, name)),
			fmt.Sprintf("./static/sheets/%s.pdf", name),
		)
		if err != nil {
			log.Fatalf("error transferring %s/sheet.pdf: %v", name, err)
		}

		_, err = bbk.CopyFile(
			path.Join(*sheetsDir, fmt.Sprintf("%s/graph.pdf", name)),
			fmt.Sprintf("./static/sheets/%s-graph.pdf", name),
		)
		if err != nil {
			log.Fatalf("error transferring %s/graph.pdf: %v", name, err)
		}
	}

	indexTemplate := template.Must(
		template.New("index.tmpl").Funcs(
			template.FuncMap{
				"title": bbk.Title,
			},
		).ParseFiles(path.Join("./static", "index.tmpl")))

	f, err := os.Create("./static/index.html")
	if err != nil {
		log.Fatalf("os.Create index.html: %v", err)
	}
	if err := indexTemplate.Execute(f, results); err != nil {
		log.Fatalf("executing index template: %v", err)
	}
	f.Close()

	f, err = os.Create("./static/results.gob")
	if err != nil {
		log.Fatalf("os.Create results.gob: %v", err)
	}
	if err := gob.NewEncoder(f).Encode(results); err != nil {
		log.Fatalf("gob encoding results: %v", err)
	}
	f.Close()
}
