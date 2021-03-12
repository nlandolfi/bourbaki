package bbk

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	namePrefix = "%!name:"
	needPrefix = "%!need:"
	mcroPrefix = "%!mcro:"
)

type ParseResult struct {
	Name        string
	Needs       []string
	Macros      []string
	Lines       []string
	NeedsParsed []*ParseResult

	// only set if returned
	//from ParseAll
	NeededBy []string

	// set in Parse
	Title string
}

func Parse(f io.Reader) *ParseResult {
	scanner := bufio.NewScanner(f)

	p := new(ParseResult)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.HasPrefix(t, namePrefix) {
			if p.Name != "" {
				log.Fatalf("%s: multiple name directives", p.Name)
			}
			p.Name = strings.TrimPrefix(t, namePrefix)
			p.Title = Title(p.Name)
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

func Title(s string) string {
	return strings.Title(strings.Join(strings.Split(s, "_"), " "))
}

func ParseAll(sheetsdir string) (map[string]*ParseResult, error) {
	files, err := ioutil.ReadDir(sheetsdir)
	if err != nil {
		return nil, err
	}
	results := map[string]*ParseResult{}
	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		sheetfile, err := os.Open(fmt.Sprintf("%s/%s/sheet.tex", sheetsdir, f.Name()))
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			log.Fatal("opening sheet.tex: %v", err)
		}

		p := Parse(sheetfile)
		if p.Name == "" {
			log.Fatalf("no name for file %v", f)
		}
		results[p.Name] = p
	}

	for _, p := range results {
		for _, n := range p.Needs {
			o, ok := results[n]
			if !ok {
				log.Fatalf("missing %s", n)
			}

			o.NeededBy = append(o.NeededBy, p.Name)
		}
	}
	return results, nil
}
