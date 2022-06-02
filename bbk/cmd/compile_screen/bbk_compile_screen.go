package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

	cmd := exec.Command("git", "rev-parse", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	gitCommit := out.String()

	// log.Printf("%d sheets", len(results))

	sheetTemplate := template.Must(
		template.New("sheet.tmpl").Funcs(
			template.FuncMap{
				"title": bbk.Title,
			},
		).ParseFiles(path.Join(*staticDir, "sheet.tmpl")))

	type SheetData struct {
		*bbk.ParseResult
		Version string
	}

	for _, p := range results {
		name := p.Name
		out, err := os.Create(filepath.Join(*staticDir, "sheets", fmt.Sprintf("%s.html", name)))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		if err := sheetTemplate.Execute(out, &SheetData{p, gitCommit[:9]}); err != nil {
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

		// BUG: if two of the same named graphics, then an error
		graphicsDir := filepath.Join(*sheetsDir, name, "graphics")
		_, err = os.Stat(graphicsDir)
		var hasGraphics bool
		if os.IsNotExist(err) {
		} else if err != nil {
		} else {
			hasGraphics = true
		}
		if hasGraphics {
			toDir := filepath.Join(*staticDir, "sheets", "graphics")
			os.Mkdir(toDir, 0700)
			fs, err := ioutil.ReadDir(graphicsDir)
			if err != nil {
				log.Fatal(err)
			}
			for _, fi := range fs {
				_, err = bbk.CopyFile(
					filepath.Join(graphicsDir, fi.Name()),
					filepath.Join(toDir, fi.Name()),
				)
				if err != nil {
					log.Printf("error transferring %s/%s: %v", p, fi.Name(), err)
				}
			}
		}
	}

	type IndexData struct {
		Results []*bbk.ParseResult
		Version string
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
	if err := indexTemplate.Execute(f, &IndexData{results, gitCommit[:9]}); err != nil {
		log.Fatalf("executing index template: %v", err)
	}
	f.Close()

	f, err = os.Create(filepath.Join(*staticDir, "results.gob"))
	if err != nil {
		log.Fatalf("os.Create results.gob: %v", err)
	}
	if err := gob.NewEncoder(f).Encode(&bbk.SearchData{results, gitCommit[:9]}); err != nil {
		log.Fatalf("gob encoding results: %v", err)
	}
	f.Close()
}
