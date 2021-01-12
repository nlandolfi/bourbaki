package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/template"
)

var (
	name = flag.String("name", "", "name of the sheet")
)

func main() {
	flag.Parse()
	if *name == "" {
		log.Fatalf("need a name for the sheet: --name <name here>")
	}

	if err := os.Mkdir(fmt.Sprintf("./sheets/%s/", *name), 0755); err != nil {
		log.Fatal(err)
	}

	mkfile, err := os.Create(fmt.Sprintf("./sheets/%s/Makefile", *name))
	if err != nil {
		log.Fatal(err)
	}
	defer mkfile.Close()

	tmpl := template.Must(template.New("").Parse(MakefileTemplate))

	if err := tmpl.Execute(mkfile, T{Name: *name, Title: nodename(*name)}); err != nil {
		log.Fatal(err)
	}

	macrosfile, err := os.Create(fmt.Sprintf("./sheets/%s/macros.tex", *name))
	if err != nil {
		log.Fatal(err)
	}
	defer macrosfile.Close()

	tmpl2 := template.Must(template.New("").Parse(MacrosTemplate))

	if err := tmpl2.Execute(macrosfile, T{Name: *name, Title: nodename(*name)}); err != nil {
		log.Fatal(err)
	}

	mainfile, err := os.Create(fmt.Sprintf("./sheets/%s/%s.tex", *name, *name))
	if err != nil {
		log.Fatal(err)
	}
	defer mainfile.Close()

	tmpl3 := template.Must(template.New("").Parse(MainTemplate))

	if err := tmpl3.Execute(mainfile, T{Name: *name, Title: nodename(*name)}); err != nil {
		log.Fatal(err)
	}

	contentfile, err := os.Create(fmt.Sprintf("./sheets/%s/content.tex", *name, *name))
	if err != nil {
		log.Fatal(err)
	}
	defer contentfile.Close()
}

type T struct {
	Name  string
	Title string
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

const MakefileTemplate = `{{ .Name }}.pdf: ../../*.tex ../../trademark.pdf *.tex
	pdflatex {{ .Name }}.tex
	make defs

make:
	make {{ .Name }}.pdf

open: {{ .Name }}.pdf
	open {{ .Name }}.pdf

o:
	make open

defs: {{ .Name }}.tex
	defs {{ .Name }}.tex

spell:
	aspell -c {{ .Name }}.tex

s:
	make spell`

const MacrosTemplate = `% {{ .Name }} macros`

const MainTemplate = `\input{{"{"}}../../sheet.tex{{"}"}}
\sbasic

\sinput{{"{"}}../{{ .Name }}/macros.tex{{"}"}}

\sstart

\stitle{{"{"}}{{ .Title }}{{"}"}}

\input{content}

\strats`
