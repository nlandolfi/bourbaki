package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"bbk"
)

var (
	sheetsDir = flag.String("sheets-dir", "../sheets", "the sheets directory")
)

func main() {
	flag.Parse()

	files, err := ioutil.ReadDir(*sheetsDir)
	if err != nil {
		log.Fatalf("ioutil.ReadDir: %q error: %v", *sheetsDir, err)
	}

	// start the entries with the header
	var entries [][]string = [][]string{
		[]string{"name", "need"},
	}

	for _, f := range files {
		sheetfile, err := os.Open(fmt.Sprintf("../sheets/%s/sheet.tex", f.Name()))

		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			log.Fatal("opening sheet.tex: %v", err)
		}

		p := bbk.Parse(sheetfile)
		for _, n := range p.Needs {
			entries = append(entries,
				[]string{
					bbk.Title(p.Name), bbk.Title(n),
				})
		}
	}

	w := csv.NewWriter(os.Stdout)

	if err := w.WriteAll(entries); err != nil {
		log.Fatal(err)
	}
}
