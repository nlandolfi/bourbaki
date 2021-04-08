package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"bbk"
)

var (
	inFile = flag.String("in", "sheet.tex", "sheet file")
	// create, makefile, graph
	mode      = flag.String("mode", "mk", "which mode {c, mk, g, ts}")
	sheetsDir = flag.String("sheets", "../", "where to find other sheets")
)

func main() {
	flag.Parse()

	results, err := bbk.ParseAll(*sheetsDir)
	if err != nil {
		log.Fatal(err)
	}

	rs := make(map[string]*bbk.ParseResult, len(results))
	for _, p := range results {
		rs[p.Name] = p
	}

	// just want to get the name
	f, err := os.Open("sheet.tex")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// just want to get the name
	g, err := os.Open("macros.tex")
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()
	partialp := bbk.Parse(f, g)
	p, ok := rs[partialp.Name]
	if !ok {
		log.Fatalf("name %q not found among sheets", partialp.Name)
	}

	switch *mode {
	case "c":
		writeFile(rs, p)
	case "g":
		writeGraph(rs, p)
	case "ts":
		w := os.Stdout
		fmt.Fprintf(w, "%d terms\n", len(p.Terms))
		for _, t := range p.Terms {
			fmt.Fprintf(w, " - %s\n", t)
		}
	case "mk":
		tmpl := template.Must(template.New("").Parse(bbk.MakefileTemplate))

		if err := tmpl.Execute(os.Stdout, p); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown mode: %q", *mode)
	}
}

func writeFile(all map[string]*bbk.ParseResult, p *bbk.ParseResult) {
	fmt.Fprintln(os.Stdout, "\\input{../../sheet.tex}")
	fmt.Fprintln(os.Stdout, "\\sbasic")
	for _, n := range p.AllNeeds {
		fmt.Fprintf(os.Stdout, "\\input{../%s/macros.tex}\n", n)
	}
	fmt.Fprintln(os.Stdout, "\\input{./macros.tex}")
	fmt.Fprintln(os.Stdout, "\\sstart")
	fmt.Fprintf(os.Stdout, "\\stitle{%s}\n", bbk.NodeName(p.Name))
	//fmt.Fprintf(os.Stdout, "{\\small Needs: %s}\n", strings.Join(p.Needs, ", "))
	for _, l := range p.Lines {
		fmt.Fprintln(os.Stdout, l)
	}
	fmt.Fprintln(os.Stdout, "\\strats")
}

func writeGraph(rs map[string]*bbk.ParseResult, p *bbk.ParseResult) {
	var rows = [][]string{
		[]string{"name", "need"},
	}
	seen := map[*bbk.ParseResult]bool{}
	q := []*bbk.ParseResult{p}
	for len(q) > 0 {
		next := q[0]
		q = q[1:]

		if seen[next] {
			continue
		}

		seen[next] = true

		for _, n := range next.Needs {
			q = append(q, rs[n])
			rows = append(rows, []string{bbk.NodeName(next.Name), bbk.NodeName(n)})
		}
	}

	// if no needs
	if len(rows) == 1 {
		// add a single entry for the node
		rows = append(rows,
			[]string{bbk.NodeName(p.Name), ""},
		)
	}
	f, err := os.Create("graph.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	if err := w.WriteAll(rows); err != nil {
		log.Fatal(err)
	}
}
