package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/template"
)

var (
	inFile    = flag.String("in", "sheet.tex", "sheet file")
	mode      = flag.String("mode", "c", "which mode {c}")
	sheetsDir = flag.String("sheets", "../", "where to find other sheets")
)

const (
	namePrefix = "%!name:"
	needPrefix = "%!need:"
	mcroPrefix = "%!mcro:"
)

type Parse struct {
	Name        string
	Needs       []string
	Macros      []string
	Lines       []string
	NeedsParsed []*Parse
}

func parse(f io.Reader) *Parse {
	scanner := bufio.NewScanner(f)

	p := new(Parse)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.HasPrefix(t, namePrefix) {
			if p.Name != "" {
				log.Fatal("multiple name directives")
			}
			p.Name = strings.TrimPrefix(t, namePrefix)
			continue
		}
		if strings.HasPrefix(t, needPrefix) {
			p.Needs = append(p.Needs, strings.TrimPrefix(t, needPrefix))
			continue
		}
		if strings.HasPrefix(t, mcroPrefix) {
			p.Macros = append(p.Macros, strings.TrimPrefix(t, mcroPrefix))
			continue
		}
		p.Lines = append(p.Lines, t)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return p
}

func parseNeeds(p *Parse) {
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
		np := parse(f)
		for _, n := range np.Needs {
			q = append(q, n)
		}
		p.NeedsParsed = append(p.NeedsParsed, np)
	}
}

func main() {
	flag.Parse()

	f, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	p := parse(f)
	parseNeeds(p)

	switch *mode {
	case "c":
		writeFile(p)
	case "mk":
		tmpl := template.Must(template.New("").Parse(MakefileTemplate))

		if err := tmpl.Execute(os.Stdout, p); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown mode: %q", *mode)
	}
}

func writeFile(p *Parse) {
	log.Print("writing file: %v", p)
	fmt.Fprintln(os.Stdout, "\\input{../../sheet.tex}")
	fmt.Fprintln(os.Stdout, "\\sbasic")
	for _, n := range p.NeedsParsed {
		fmt.Fprintf(os.Stdout, "\\input{../%s/macros.tex}\n", n.Name)
	}
	fmt.Fprintln(os.Stdout, "\\input{./macros.tex}")
	fmt.Fprintln(os.Stdout, "\\sstart")
	fmt.Fprintf(os.Stdout, "\\stitle{%s}\n", nodename(p.Name))
	for _, l := range p.Lines {
		fmt.Fprintln(os.Stdout, l)
	}
	fmt.Fprintln(os.Stdout, "\\strats")
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

const MakefileTemplate = `{{ .Name }}.tex: sheet.tex
	sheet -mode c -in sheet.tex > {{ .Name }}.tex

{{ .Name }}.pdf: ../../*.tex ../../trademark.pdf *.tex {{ .Name }}.tex
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
