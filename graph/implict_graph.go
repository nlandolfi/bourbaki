package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("../sheets/")
	if err != nil {
		log.Fatalf("ioutil.ReadDir: error: %v", err)
	}

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

		p := parse(sheetfile)
		for _, n := range p.Needs {
			entries = append(entries,
				[]string{
					title(p.Name), title(n),
				})
		}
	}

	w := csv.NewWriter(os.Stdout)

	if err := w.WriteAll(entries); err != nil {
		log.Fatal(err)
	}
}

const (
	namePrefix = "%!name:"
	needPrefix = "%!need:"
	mcroPrefix = "%!mcro:"
)

type Parse struct {
	Name        string
	Needs       []string
	Macros      []string
	Lines       []string
	NeedsParsed []*Parse
}

func parse(f io.Reader) *Parse {
	scanner := bufio.NewScanner(f)

	p := new(Parse)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.HasPrefix(t, namePrefix) {
			if p.Name != "" {
				log.Fatalf("%s: multiple name directives", p.Name)
			}
			p.Name = strings.TrimPrefix(t, namePrefix)
			continue
		}
		if strings.HasPrefix(t, needPrefix) {
			p.Needs = append(p.Needs, strings.TrimPrefix(t, needPrefix))
			continue
		}
		if strings.HasPrefix(t, mcroPrefix) {
			p.Macros = append(p.Macros, strings.TrimPrefix(t, mcroPrefix))
			continue
		}
		p.Lines = append(p.Lines, t)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return p
}

func title(s string) string {
	return strings.Title(strings.Join(strings.Split(s, "_"), " "))
}
