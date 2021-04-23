package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"bbk"
)

var (
	sheetsDir = flag.String("sheets-dir", "../sheets", "the sheets directory")
	staticDir = flag.String("static-dir", "./static", "directory of static files")
)

func main() {
	flag.Parse()

	results, err := bbk.ParseAll(*sheetsDir)
	if err != nil {
		log.Fatal(err)
	}

	if len(results) == 0 {
		log.Printf("warning: no sheets found")
	}

	// log.Printf("%d sheets", len(results))

	sheetTemplate := template.Must(
		template.New("sheet.tmpl").Funcs(
			template.FuncMap{
				"title": bbk.Title,
			},
		).ParseFiles(path.Join(*staticDir, "sheet.tmpl")))

	for _, p := range results {
		name := p.Name
		out, err := os.Create(filepath.Join(*staticDir, "sheets", fmt.Sprintf("%s.html", name)))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		if err := sheetTemplate.Execute(out, p); err != nil {
			log.Fatal(err)
		}
		out.Close()

		_, err = bbk.CopyFile(
			filepath.Join(*sheetsDir, name, fmt.Sprintf("%s.pdf", name)),
			filepath.Join(*staticDir, "sheets", fmt.Sprintf("%s.pdf", name)),
		)
		if err != nil {
			log.Printf("error transferring %s/sheet.pdf: %v", name, err)
		}

		_, err = bbk.CopyFile(
			filepath.Join(*sheetsDir, name, "graph.pdf"),
			filepath.Join(*staticDir, "sheets", fmt.Sprintf("%s-graph.pdf", name)),
		)
		if err != nil {
			log.Printf("error transferring %s/graph.pdf: %v", name, err)
		}
	}

	indexTemplate := template.Must(
		template.New("index.tmpl").Funcs(
			template.FuncMap{
				"title": bbk.Title,
			},
		).ParseFiles(filepath.Join(*staticDir, "index.tmpl")))

	f, err := os.Create(filepath.Join(*staticDir, "index.html"))
	if err != nil {
		log.Fatalf("os.Create index.html: %v", err)
	}
	if err := indexTemplate.Execute(f, results); err != nil {
		log.Fatalf("executing index template: %v", err)
	}
	f.Close()

	f, err = os.Create(filepath.Join(*staticDir, "results.gob"))
	if err != nil {
		log.Fatalf("os.Create results.gob: %v", err)
	}
	if err := gob.NewEncoder(f).Encode(results); err != nil {
		log.Fatalf("gob encoding results: %v", err)
	}
	f.Close()
}
