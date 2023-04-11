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
		if !f.IsDir() {
			continue
		}

		litfile, err := os.Open(fmt.Sprintf("../sheets/%s/sheet.lit", f.Name()))
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			log.Fatalf("opening sheet.lit: %v", err)
		}

		p, err := bbk.ParseSheet(litfile)
		if p.Config.Name == "" {
			log.Fatalf("no name for file %v", f)
		}
		for _, n := range p.Config.Needs {
			entries = append(entries,
				[]string{
					bbk.Title(p.Config.Name), bbk.Title(n),
				})
		}
		// add leaf entries
		if len(p.Config.Needs) == 0 {
			entries = append(entries, []string{bbk.Title(p.Config.Name), ""})
		}

		if err := litfile.Close(); err != nil {
			log.Fatal(err)
		}
	}

	w := csv.NewWriter(os.Stdout)

	if err := w.WriteAll(entries); err != nil {
		log.Fatal(err)
	}
}
