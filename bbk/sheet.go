package bbk

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
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
	Body        string
	NeedsParsed []*ParseResult
	Terms       []string

	// only set if returned
	//from ParseAll
	NeededBy []string
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
	p.Body = strings.Join(p.Lines, " ")
	p.Terms = Terms(p.Body)
	return p
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
			log.Printf("%q is a sheets directory, but is missing sheets.tex", f.Name())
			continue
		} else if err != nil {
			log.Fatalf("opening sheet.tex: %v", err)
		}

		p := Parse(sheetfile)
		if p.Name == "" {
			log.Fatalf("no name for file %v", f)
		}
		results[p.Name] = p
		sheetfile.Close()
	}

	for _, p := range results {
		for _, n := range p.Needs {
			o, ok := results[n]
			if !ok {
				log.Fatalf("%s refers to %s, which is missing", p.Name, n)
			}

			o.NeededBy = append(o.NeededBy, p.Name)
		}
	}

	for _, p := range results {
		sort.Strings(p.Needs)
		sort.Strings(p.NeededBy)
	}
	return results, nil
}

func Terms(s string) []string {
	ts := make([]string, 0)
	for {
		i := strings.Index(s, "\\ct{")
		if i == -1 {
			break
		}
		s = s[i+4:]
		j := strings.Index(s, "}")
		ts = append(ts, s[:j])
	}
	return ts
}
