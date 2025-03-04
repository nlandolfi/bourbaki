package main

import (
	"bbk"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	graphFile = flag.String("graph-file", "graph.csv", "the graph to check")
)

func main() {
	flag.Parse()

	f, err := os.Open(*graphFile)
	if err != nil {
		log.Fatalf("os.Open: %q, %v", *graphFile, err)
	}
	defer f.Close()

	entries, err := bbk.ParseGraph(f)
	if err != nil {
		log.Fatal(err)
	}

	names := []string{}
	for _, e := range entries {
		_, err := os.Stat(fmt.Sprintf("../sheets/%s", bbk.DirName(e.Name)))
		if os.IsNotExist(err) {
			log.Fatalf("sheets/%s does not exist", bbk.DirName(e.Name))
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
		// if no needs
		if len(rows) == 1 {
			// add a single entry for the node
			rows = append(rows,
				[]string{e.Name, ""},
			)
		}

		written := fmt.Sprintf("./clips/%s.csv", bbk.DirName(e.Name))
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

	fmt.Fprintf(f, "all: ../graph.tmpl\n")
	for _, s := range names {
		d := bbk.DirName(s)
		fmt.Fprintf(f, "\tbbk_graph --csv %s.csv --tmpl ../graph.tmpl > %s.graphviz\n", d, d)
		fmt.Fprintf(f, "\tdot %s.graphviz -o %s.pdf -T pdf\n", d, d)
	}
}
