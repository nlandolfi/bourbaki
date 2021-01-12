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
		out, err := os.Create(fmt.Sprintf("./sheets/%s.html", dirname(s)))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		needs := make(map[string]string)
		if e, ok := entries[s]; ok {
			for _, ss := range e.Needs {
				needs[ss] = dirname(ss)
			}
		}
		compile2(s, needs, f, out)
		compiled[s] = &CompiledInfo{
			DirName: dirname(s),
			Needs:   needs,
		}

		out.Close()
		f.Close()
	}

	f, err = os.Create("./index.html")
	if err != nil {
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

const HTMLTemplate = `<!DOCTYPE html>
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

const HTMLTemplate2 = `<!DOCTYPE html>
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
		<a href="../index.html">Back to index</a>
		</div>

		<iframe id="sheet" src="../../../sheets/{{ .DirName }}/{{ .DirName }}.pdf">
		</iframe>

		<div class="info">
		<a href="../../../graph/clips/{{ .DirName }}.pdf"> See graph on own page </a>
		</div>
		<iframe id="graph" src="../../../graph/clips/{{ .DirName }}.pdf">
		</iframe>
  </body>
</html>`

func compile2(s string, needs map[string]string, f *os.File, into io.Writer) {
	if err := tmpl2.Execute(into, Data{
		Needs:   needs,
		DirName: dirname(s),
	}); err != nil {
		log.Fatal(err)
	}
}

func compile(s string, needs map[string]string, f *os.File, into io.Writer) {
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

	// the order of these matters,
	// since the second will match all of the enumerate block
	re19 := regexp.MustCompile(`\\item((?s).+)\\item`)
	contents = re19.ReplaceAll(contents, []byte(`<li>$1</li>\item`))
	re18 := regexp.MustCompile(`\\item((?s).+)\\end\{enumerate\}`)
	contents = re18.ReplaceAll(contents, []byte(`<li>$1</li>\end{enumerate}`))
	re20 := regexp.MustCompile(`\\item((?s).+)\\end\{itemize\}`)
	contents = re20.ReplaceAll(contents, []byte(`<li>$1</li>\end{itemize}`))

	re16 := regexp.MustCompile(`\\begin\{enumerate\}\n`)
	contents = re16.ReplaceAll(contents, []byte(`<ol>`))
	re17 := regexp.MustCompile(`\\end\{enumerate\}\n`)
	contents = re17.ReplaceAll(contents, []byte(`</ol>`))

	reBeginProp := regexp.MustCompile(`\\begin\{prop\}\n`)
	contents = reBeginProp.ReplaceAll(contents, []byte(`<div class="prop"><b> Proposition </b>`))
	reEndProp := regexp.MustCompile(`\\end\{prop\}\n`)
	contents = reEndProp.ReplaceAll(contents, []byte(`</div>`))

	reBeginThm := regexp.MustCompile(`\\begin\{thm\}\n`)
	contents = reBeginThm.ReplaceAll(contents, []byte(`<div class="thm"><b> Theorem </b>`))
	reEndThm := regexp.MustCompile(`\\end\{thm\}\n`)
	contents = reEndThm.ReplaceAll(contents, []byte(`</div>`))

	reBeginDefn := regexp.MustCompile(`\\begin\{defn\}\n`)
	contents = reBeginDefn.ReplaceAll(contents, []byte(`<div class="defn"><b> Definition </b>`))
	reEndDefn := regexp.MustCompile(`\\end\{defn\}\n`)
	contents = reEndDefn.ReplaceAll(contents, []byte(`</div>`))

	reBeginExpl := regexp.MustCompile(`\\begin\{expl\}\n`)
	contents = reBeginExpl.ReplaceAll(contents, []byte(`<div class="expl"><b> Example </b>`))
	reEndExpl := regexp.MustCompile(`\\end\{expl\}\n`)
	contents = reEndExpl.ReplaceAll(contents, []byte(`</div>`))

	reBeginProof := regexp.MustCompile(`\\begin\{proof\}\n`)
	contents = reBeginProof.ReplaceAll(contents, []byte(`<div class="expl"><i> Proof. </i>`))
	reEndProof := regexp.MustCompile(`\\end\{proof\}\n`)
	contents = reEndProof.ReplaceAll(contents, []byte(`</div>`))

	reBeginItemize := regexp.MustCompile(`\\begin\{itemize\}\n`)
	contents = reBeginItemize.ReplaceAll(contents, []byte(`<ul>`))
	reEndItemize := regexp.MustCompile(`\\end\{itemize\}\n`)
	contents = reEndItemize.ReplaceAll(contents, []byte(`</ul>`))

	// paragraphs
	re14 := regexp.MustCompile(`\n\n([A-z])`)
	re15 := regexp.MustCompile(`([^>\n])(\n)\n`)
	contents = re14.ReplaceAll(contents, []byte(`<p>$1`))
	contents = re15.ReplaceAll(contents, []byte(`$1</p>`))

	if err := tmpl.Execute(into, Data{
		DirName: dirname(s),
		Content: contents,
		Needs:   needs,
	}); err != nil {
		log.Fatal(err)
	}
}

var tmpl = template.Must(template.New("").Parse(HTMLTemplate))
var tmpl2 = template.Must(template.New("").Parse(HTMLTemplate2))

type Data struct {
	DirName string
	Content []byte
	Needs   map[string]string
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
	<img src="../../trademark.pdf" id="trademark"><h1>HyperText Index</h1>
<ul>
{{ range $k, $v := .Compiled }}
	<li> <a href="./sheets/{{ $v.DirName }}.html"> {{ $k }} </a> <br> <a href="../../graph/clips/{{ $v.DirName}}.pdf"> Graph </a> <br> (Needs: {{ range $k, $v := .Needs }}<a href="./sheets/{{ $v }}.html">{{ $k }};</a>{{ end }})</li>
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
