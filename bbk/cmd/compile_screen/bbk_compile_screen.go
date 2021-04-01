package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"text/template"

	"bbk"
)

var (
	sheetsDir = flag.String("sheets-dir", "../../sheets", "the sheets directory")
)

func main() {
	flag.Parse()

	results, err := bbk.ParseAll(*sheetsDir)
	if err != nil {
		log.Fatal(err)
	}

	names := make([]string, 0, len(results))
	for n := range results {
		names = append(names, n)
	}
	sort.Strings(names)

	log.Printf("%d sheets", len(results))

	for _, name := range names {
		p := results[name]
		out, err := os.Create(fmt.Sprintf("./static/sheets/%s.html", name))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		if err := sheetTemplate.Execute(out, p); err != nil {
			log.Fatal(err)
		}
		out.Close()

		_, err = bbk.CopyFile(
			fmt.Sprintf("../../sheets/%s/%s.pdf", name, name),
			fmt.Sprintf("./static/sheets/%s.pdf", name),
		)
		if err != nil {
			log.Fatalf("error transferring %s/sheet.pdf: %v", name, err)
		}

		_, err = bbk.CopyFile(
			fmt.Sprintf("../../sheets/%s/graph.pdf", name),
			fmt.Sprintf("./static/sheets/%s_graph.pdf", name),
		)
		if err != nil {
			log.Fatalf("error transferring %s/graph.pdf: %v", name, err)
		}
	}

	f, err := os.Create("./static/index.html")
	if err != nil {
		log.Fatalf("os.Create index.html: %v", err)
	}
	if err := indexTemplate.Execute(f, results); err != nil {
		log.Fatalf("executing index template: %v", err)
	}
	f.Close()

	f, err = os.Create("./static/results.gob")
	if err != nil {
		log.Fatalf("os.Create results.gob: %v", err)
	}
	if err := gob.NewEncoder(f).Encode(results); err != nil {
		log.Fatalf("gob encoding results: %v", err)
	}
	f.Close()
}

const IndexTemplate = `<!DOCTYPE html>
  <head>
    <link rel="stylesheet" href="./style.css">
    <link rel="stylesheet" href="./fonts.css">
	</head>
  <body>
		<div class="page">
			<div class="content">
				<img src="../trademark.pdf" id="trademark">
				<h1>HyperText Index</h1>
				<form class="search" action="/search" method="GET" autofocus>
				<input class="search" placeholder="search bourbaki..." name="query">
				</form>
				<br>
				<a href="./sheets/introduction.html" style="margin:0px auto">View the project introduction.</a>
				<ul>
				{{ range $k, $v := . }}<li> <a href="./sheets/{{ $v.Name }}.html">{{ title .Name }}</a> </li>{{ end }}
				</ul>
		</div>
	</div>
	</body>
</html>
`

const SheetTemplate = `<!DOCTYPE html>
  <head>
    <link rel="stylesheet" href="../style.css">
    <link rel="stylesheet" href="../fonts.css">
  </head>
  <body>
		<div class="info">
		{{ if .Needs }}
		Needs:
		<ul>
			{{ range $k, $v := .Needs }}
				<li> <a tabindex="0" href="./{{ $v }}.html"> {{ title $v }} </a> </li>
			{{ end }}
		</ul>
		{{ else }}
		No needs.
		{{ end }}
		{{ if .NeededBy }}
		Needed by:
		<ul>
			{{ range $k, $v := .NeededBy }}
				<li> <a tabindex="0" href="./{{ $v }}.html"> {{ title $v }} </a> </li>
			{{ end }}
		</ul>
		{{ else }}
		{{ end }}
		<form class="search" action="/search" method="GET">
		<input class="search"placeholder="Search" name="query">
		<a tabindex="0" href="../index.html">Back to index</a>
		</form>
		<br/>
		<br/>
		<a tabindex="0" href="./{{ .Name }}.pdf"> See sheet on own page </a>
		</div>

		<iframe id="sheet" src="./{{ .Name }}.pdf">
		</iframe>

		<div class="info">
		<a tabindex="0" href="./{{ .Name }}_graph.pdf"> See graph on own page </a>
		</div>
		<iframe id="graph" src="./{{ .Name }}_graph.pdf">
		</iframe>
  </body>
</html>`

var sheetTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	"title": bbk.Title,
}).Parse(SheetTemplate))
var indexTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	"title": bbk.Title,
}).Parse(IndexTemplate))
