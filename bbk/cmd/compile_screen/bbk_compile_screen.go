package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"bbk"
)

type Entry struct {
	Name     string
	Needs    []string
	NeededBy []string
}

const (
	sheetsFormat = "./static/sheets/%s.html"
)

var (
	graphCSVFile = flag.String("graph-csv", "../../graph/graph.csv", "csv of the graph; probably implicit")
)

func main() {
	flag.Parse()
	f, err := os.Open(*graphCSVFile)
	if err != nil {
		log.Fatal("os.Open: %v", err)
	}
	defer f.Close()
	entries, err := bbk.ParseGraph(f)
	if err != nil {
		log.Fatal(err)
	}

	seen := make(map[string]bool)
	for _, e := range entries {
		seen[e.Name] = true
		for _, n := range e.Needs {
			seen[n] = true
		}
	}

	compiled := make(map[string]*CompiledInfo)

Seen:
	for s := range seen {
		f, err := os.Open(fmt.Sprintf("../../sheets/%s/sheet.tex", bbk.DirName(s)))
		if os.IsNotExist(err) {
			log.Printf("missing: %s", bbk.DirName(s))
			continue Seen
		} else if err != nil {
			log.Fatalf("os.Open: %v", err)
		}
		out, err := os.Create(fmt.Sprintf("./static/sheets/%s.html", bbk.DirName(s)))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		needs := make(map[string]string)
		if e, ok := entries[s]; ok {
			for _, ss := range e.Needs {
				needs[ss] = bbk.DirName(ss)
			}
		}
		needed_by := make(map[string]string)
		if e, ok := entries[s]; ok {
			for _, ss := range e.NeededBy {
				needed_by[ss] = bbk.DirName(ss)
			}
		}
		compile(s, needs, needed_by, f, out)
		compiled[s] = &CompiledInfo{
			DirName:  bbk.DirName(s),
			Needs:    needs,
			NeededBy: needed_by,
		}

		_, err = bbk.CopyFile(
			fmt.Sprintf("../../sheets/%s/%s.pdf", bbk.DirName(s), bbk.DirName(s)),
			fmt.Sprintf("./static/sheets/%s.pdf", bbk.DirName(s)),
		)
		if err != nil {
			log.Printf("error: %v", err)
		}

		_, err = bbk.CopyFile(
			fmt.Sprintf("../../graph/clips/%s.pdf", bbk.DirName(s)),
			fmt.Sprintf("./static/sheets/%s_graph.pdf", bbk.DirName(s)),
		)
		if err != nil {
			log.Printf("error: %v", err)
		}

		out.Close()
		f.Close()
	}

	f, err = os.Create("./static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := indexTemplate.Execute(f, IndexData{compiled}); err != nil {
		log.Fatal(err)
	}
}

func compile(s string, needs, needed_by map[string]string, f *os.File, into io.Writer) {
	if err := sheetTemplate.Execute(into, Data{
		Needs:    needs,
		DirName:  bbk.DirName(s),
		NeededBy: needed_by,
	}); err != nil {
		log.Fatal(err)
	}
}

type Data struct {
	DirName  string
	Content  []byte
	Needs    map[string]string
	NeededBy map[string]string
}

const IndexTemplate = `<!DOCTYPE html>
  <head>
	  <head>
    <link rel="stylesheet" href="./style.css">
    <link rel="stylesheet" href="./fonts.css">

		<link rel="stylesheet" href="./katex/katex.min.css">
		<script defer src="./katex/katex.min.js"></script>
		<script defer src="./katex/auto-render.min.js"
						onload="renderMathInElement(document.body, { delimiters: [
						  {left: "$$", right: "$$", display: true},
							  {left: "$", right: "$", display: false},
								  {left: "\\(", right: "\\)", display: false},
									  {left: "\\[", right: "\\]", display: true}
										]});"></script>
  </head>
  <body>
	<div class="page">
	<div class="content">
	<img src="../trademark.pdf" id="trademark"><h1>HyperText Index</h1>
<ul>
{{ range $k, $v := .Compiled }}
	<li> <a href="./sheets/{{ $v.DirName }}.html"> {{ $k }} </a> <br> <a href="../../../graph/clips/{{ $v.DirName}}.pdf"> Graph </a> <br> (Needs: {{ range $k, $v := .Needs }}<a href="./sheets/{{ $v }}.html">{{ $k }};</a>{{ end }})</li>
{{ end }}
</ul>
</div>
</div>
</body>
</html>
`

type IndexData struct {
	Compiled map[string]*CompiledInfo
}

type CompiledInfo struct {
	Contents []byte
	DirName  string
	Needs    map[string]string
	NeededBy map[string]string
}

const SheetTemplate = `<!DOCTYPE html>
  <head>
    <link rel="stylesheet" href="../style.css">
    <link rel="stylesheet" href="../fonts.css">
		<link rel="stylesheet" href="../katex/katex.min.css">
		<script defer src="../katex/katex.min.js"></script>
		<script defer src="../katex/auto-render.min.js"
						onload="renderMathInElement(document.body, {delimiters: [
							{left: '$$', right: '$$', display: true},
							{left: '$', right: '$', display: false},
							{left: '\\(', right: '\\)', display: false},
							{left: '\\[', right: '\\]', display: true}]});"></script>
  </head>
  <body>
		<div class="info">
		{{ if .Needs }}
		Needs:
		<ul>
			{{ range $k, $v := .Needs }}
				<li> <a href="./{{ $v }}.html"> {{ $k }} </a> </li>
			{{ end }}
		</ul>
		{{ else }}
		No needs.
		{{ end }}
		{{ if .NeededBy }}
		Needed by:
		<ul>
			{{ range $k, $v := .NeededBy }}
				<li> <a href="./{{ $v }}.html"> {{ $k }} </a> </li>
			{{ end }}
		</ul>
		{{ else }}
		{{ end }}
		<a href="../index.html">Back to index</a>
		</div>

		<iframe id="sheet" src="./{{ .DirName }}.pdf">
		</iframe>

		<div class="info">
		<a href="./{{ .DirName }}_clip.pdf"> See graph on own page </a>
		</div>
		<iframe id="graph" src="./{{ .DirName }}_clip.pdf">
		</iframe>
  </body>
</html>`

var sheetTemplate = template.Must(template.New("").Parse(SheetTemplate))
var indexTemplate = template.Must(template.New("").Parse(IndexTemplate))
