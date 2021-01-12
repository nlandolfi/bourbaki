package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Entry struct {
	Name  string
	Needs []string
}

func main() {
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

	seen := make(map[string]bool)
	for _, e := range entries {
		seen[e.Name] = true
		for _, n := range e.Needs {
			seen[n] = true
		}
	}

	names := []string{}
	for _, e := range entries {
		_, err := os.Stat(fmt.Sprintf("../sheets/%s", dirname(e.Name)))
		if os.IsNotExist(err) {
			continue
			log.Print(e.Name)
		} else if err != nil {
			log.Fatal(err)
		}

		// graph search
		var seen = map[string]bool{}
		var q []string = []string{e.Name}
		var rows = [][]string{
			[]string{"name", "need"},
		}
		for len(q) > 0 {
			next := q[0]
			q = q[1:]

			if seen[next] {
				continue
			}

			seen[next] = true

			te, ok := entries[next]
			if !ok {
				continue
			}

			for _, n := range te.Needs {
				q = append(q, n)
				rows = append(rows, []string{next, n})
			}
		}

		written := fmt.Sprintf("./clips/%s.csv", dirname(e.Name))
		f, err := os.Create(written)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w := csv.NewWriter(f)
		if err := w.WriteAll(rows); err != nil {
			log.Fatal(err)
		}

		names = append(names, e.Name)
	}

	f, err = os.Create("./clips/Makefile")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "all: ../graph.tmpl ../clip.go \n")
	for _, s := range names {
		d := dirname(s)
		fmt.Fprintf(f, "\tgraph --csv %s.csv --tmpl ../graph.tmpl > %s.graphviz\n", d, d)
		fmt.Fprintf(f, "\tdot %s.graphviz -o %s.pdf -T pdf\n", d, d)
	}
}

func nodename(s string) string {
	pieces := strings.Split(s, "_")
	return strings.Title(strings.Join(pieces, " "))
}

func dirname(s string) string {
	pieces := strings.Split(s, " ")
	for i, p := range pieces {
		pieces[i] = strings.ToLower(p)
	}
	return strings.Join(pieces, "_")
}
