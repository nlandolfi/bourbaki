package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/alecthomas/template"
)

type Entry struct {
	Name  string
	Needs []string
}

func main() {
	f, err := os.Open("../../graph/graph.csv")
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

	compiled := make(map[string]*CompiledInfo)

Seen:
	for s := range seen {
		f, err := os.Open(fmt.Sprintf("../../sheets/%s/%s.tex", dirname(s), dirname(s)))
		if os.IsNotExist(err) {
			log.Printf("missing: %s", dirname(s))
			continue Seen
		} else if err != nil {
			log.Fatalf("os.Open: %v", err)
		}
		out, err := os.Open(fmt.Sprintf("./sheets/%s.html", dirname(s)))
		if os.IsNotExist(err) {
			out, err = os.Create(fmt.Sprintf("./sheets/%s.html", dirname(s)))
			if err != nil {
				log.Fatalf("os.Create: %v", err)
			}
			if err := out.Truncate(0); err != nil {
				log.Fatal(err)
			}
		} else if err != nil {
			log.Fatalf("os.Open: %v", err)
		}
		needs := make(map[string]string)
		if e, ok := entries[s]; ok {
			for _, ss := range e.Needs {
				needs[ss] = dirname(ss)
			}
		}
		compile(needs, f, out)
		compiled[s] = &CompiledInfo{
			DirName: dirname(s),
			Needs:   needs,
		}

		out.Close()
		f.Close()
	}

	f, err = os.Open("./index.html")
	if os.IsNotExist(err) {
		f, err = os.Create("./index.html")
		if err != nil {
			log.Fatal(err)
		}
		if err := f.Truncate(0); err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	tmpl2 := template.Must(template.New("").Parse(IndexTemplate))
	if err := tmpl2.Execute(f, IndexData{compiled}); err != nil {
		log.Fatal(err)
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

const HTMLTemplate = `<html>
  <head>
    <link rel="stylesheet" href="../style.css">
    <link rel="stylesheet" href="../fonts.css">
  </head>
  <body>
		{{ if .Needs }}
		<div class="info">
		Needs:
		<ul>
			{{ range $k, $v := .Needs }}
				<li> <a href="./{{ $v }}.html"> {{ $k }} </a> </li>
			{{ end }}
		</ul>
		</div>
		{{ else }}
		<div class="info">
		No needs. <a href="../index.html">Back to index</a>
		</div>
		{{ end }}
    <div class="page">
		{{ printf "%s" .Content }}
		</div>
  </body>
</html>`

func compile(needs map[string]string, f *os.File, into io.Writer) {
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("ioutil.ReadAll: %v", err)
	}
	re1 := regexp.MustCompile(`\\ct\{(.+)\}\{.*\}`)
	re2 := regexp.MustCompile(`\\textit\{(.+)\}`)
	re3 := regexp.MustCompile(`\\textbf\{(.+)\}`)
	re4 := regexp.MustCompile(`\\say\{(.+)\}`)
	re5 := regexp.MustCompile(`\%(.+)\n`)
	re6 := regexp.MustCompile(`\\stitle\{(.+)\}`)
	re7 := regexp.MustCompile(`\\ssection\{(.+)\}`)
	re8 := regexp.MustCompile(`\\ssubsection\{(.+)\}`)
	re9 := regexp.MustCompile(`\\sstart\n`)
	re10 := regexp.MustCompile(`\\strats\n`)
	re11 := regexp.MustCompile(`\\sbasic\n`)
	re12 := regexp.MustCompile(`\\input\{(.+)\}\n`)
	re13 := regexp.MustCompile(`\\sinput\{(.+)\}\n`)

	// strip comments
	contents = re5.ReplaceAll(contents, []byte(``))
	// \ct, \textit, \textbf
	contents = re1.ReplaceAll(contents, []byte("<b>$1</b>"))
	contents = re2.ReplaceAll(contents, []byte("<i>$1</i>"))
	contents = re3.ReplaceAll(contents, []byte("<b>$1</b>"))

	// \say
	contents = re4.ReplaceAll(contents, []byte(`"$1"`))

	// \stitle, \ssection, \ssubsection
	contents = re6.ReplaceAll(contents, []byte(`<img src="../../trademark.pdf" id="trademark"><h1>$1</h1>`))
	contents = re7.ReplaceAll(contents, []byte(`<h2>$1</h1>`))
	contents = re8.ReplaceAll(contents, []byte(`<h3>$1</h3>`))

	// \sstart \strats
	contents = re9.ReplaceAll(contents, []byte(`<div class="content">
`))
	contents = re10.ReplaceAll(contents, []byte(`</div>
`))

	// \sbasic
	contents = re11.ReplaceAll(contents, []byte(``))
	// \input
	contents = re12.ReplaceAll(contents, []byte(``))
	// \sinput
	contents = re13.ReplaceAll(contents, []byte(``))

	// paragraphs
	re14 := regexp.MustCompile(`\n\n([A-z])`)
	re15 := regexp.MustCompile(`([^>\n])(\n)\n`)
	contents = re14.ReplaceAll(contents, []byte(`<p>$1`))
	contents = re15.ReplaceAll(contents, []byte(`$1</p>`))

	if err := tmpl.Execute(into, Data{
		Content: contents,
		Needs:   needs,
	}); err != nil {
		log.Fatal(err)
	}
}

var tmpl = template.Must(template.New("").Parse(HTMLTemplate))

type Data struct {
	Content []byte
	Needs   map[string]string
}

const IndexTemplate = `
<html>
  <head>
    <link rel="stylesheet" href="./style.css">
    <link rel="stylesheet" href="./fonts.css">
  </head>
  <body>
	<div class="page">
	<div class="content">
	<img src="../../trademark.pdf" id="trademark"><h1>HyperText Index</h1>
<ul>
{{ range $k, $v := .Compiled }}
	<li> <a href="./sheets/{{ $v.DirName }}.html"> {{ $k }} </a> <br> (Needs: {{ range $k, $v := .Needs }}<a href="./sheets/{{ $v }}.html">{{ $k }};</a>{{ end }})</li>
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
}
