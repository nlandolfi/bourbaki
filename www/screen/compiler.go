package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/alecthomas/template"
)

var (
	inFile  = flag.String("in", "in.tex", "file to compile")
	outFile = flag.String("out", "out.html", "file to write to")
)

const HTMLTemplate = `<html>
  <head>
    <link rel="stylesheet" href="style.css">
    <link rel="stylesheet" href="fonts.css">
  </head>
  <body>
    <div class="page">
		{{ printf "%s" .Content }}
  </div>
  </body>
</html>`

func main() {
	flag.Parse()
	f, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
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
	contents = re6.ReplaceAll(contents, []byte(`<img src="../trademark.pdf" id="trademark"><h1>$1</h1>`))
	contents = re7.ReplaceAll(contents, []byte(`<h2>$1</h1>`))
	contents = re8.ReplaceAll(contents, []byte(`<h3>$1</h3>`))

	// \sstart \strats
	contents = re9.ReplaceAll(contents, []byte(`<div class="content">
`))
	contents = re10.ReplaceAll(contents, []byte(`<div>
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

	tmpl := template.Must(template.New("").Parse(HTMLTemplate))

	if err := tmpl.Execute(os.Stdout, Data{Content: contents}); err != nil {
		log.Fatal(err)
	}
}

type Data struct {
	Content []byte
}
