package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"bbk"
)

var (
	inFile = flag.String("in", "sheet.tex", "sheet file")
	// create, makefile, graph
	mode      = flag.String("mode", "mk", "which mode {c, mk, g}")
	sheetsDir = flag.String("sheets", "../", "where to find other sheets")
)

func parseNeeds(p *bbk.ParseResult) {
	var q = make([]string, len(p.Needs))
	copy(q, p.Needs)
	seen := map[string]bool{}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if seen[n] {
			continue
		}
		seen[n] = true

		fn := fmt.Sprintf("%s%s/sheet.tex", *sheetsDir, n)
		f, err := os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Parsing need: %s", n)
		np := bbk.Parse(f)
		for _, n := range np.Needs {
			q = append(q, n)
		}
		p.NeedsParsed = append(p.NeedsParsed, np)
	}
}

func main() {
	flag.Parse()

	rs, _ := bbk.ParseAll(*sheetsDir)
	wd, _ := os.Getwd()
	name := path.Base(wd)
	p := rs[name]

	switch *mode {
	case "c":
		writeFile(p)
	case "g":
		writeGraph(rs, p)
	case "mk":
		tmpl := template.Must(template.New("").Parse(MakefileTemplate))

		if err := tmpl.Execute(os.Stdout, p); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown mode: %q", *mode)
	}
}

func writeFile(p *bbk.ParseResult) {
	log.Printf("writing file: %v", p)
	fmt.Fprintln(os.Stdout, "\\input{../../sheet.tex}")
	fmt.Fprintln(os.Stdout, "\\sbasic")
	for _, n := range p.NeedsParsed {
		fmt.Fprintf(os.Stdout, "\\input{../%s/macros.tex}\n", n.Name)
	}
	fmt.Fprintln(os.Stdout, "\\input{./macros.tex}")
	fmt.Fprintln(os.Stdout, "\\sstart")
	fmt.Fprintf(os.Stdout, "\\stitle{%s}\n", bbk.NodeName(p.Name))
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

const MakefileTemplate = `# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make {{ .Name }}.pdf
	make graph.pdf

{{ .Name }}.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > {{ .Name }}.tex

{{ .Name }}.pdf: ../../*.tex ../../trademark.pdf *.tex {{ .Name }}.tex
	pdflatex {{ .Name }}.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: {{ .Name }}.pdf
	open {{ .Name }}.pdf

o:
	make open

defs: {{ .Name }}.tex
	defs {{ .Name }}.tex

spell:
	aspell -c {{ .Name }}.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm {{.Name}}.*
	make reset
	bbk_sheet -mode mk > Makefile`
