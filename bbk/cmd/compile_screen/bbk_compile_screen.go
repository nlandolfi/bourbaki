package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	log.Printf("%d sheets", len(results))

	for name, p := range results {
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
				<a href="./sheets/introduction.html" style="margin:0px auto">View the project introduction.</a>
				<ul>
				{{ range $k, $v := . }}<li> <a href="./sheets/{{ $v.Name }}.html">{{ .Title }}</a> </li>{{ end }}
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
				<li> <a href="./{{ $v }}.html"> {{ title $v }} </a> </li>
			{{ end }}
		</ul>
		{{ else }}
		No needs.
		{{ end }}
		{{ if .NeededBy }}
		Needed by:
		<ul>
			{{ range $k, $v := .NeededBy }}
				<li> <a href="./{{ $v }}.html"> {{ title $v }} </a> </li>
			{{ end }}
		</ul>
		{{ else }}
		{{ end }}
		<a href="../index.html">Back to index</a>
		<br>
		<a href="./{{ .Name }}.pdf"> See sheet on own page </a>
		</div>

		<iframe id="sheet" src="./{{ .Name }}.pdf">
		</iframe>

		<div class="info">
		<a href="./{{ .Name }}_graph.pdf"> See graph on own page </a>
		</div>
		<iframe id="graph" src="./{{ .Name }}_graph.pdf">
		</iframe>
  </body>
</html>`

var sheetTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	"title": bbk.Title,
}).Parse(SheetTemplate))
var indexTemplate = template.Must(template.New("").Parse(IndexTemplate))
