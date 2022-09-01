package bbk

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/nlandolfi/lit"
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
	AST         *SheetAST
	MacrosLines []string
	AllNeeds    []string

	// only set if returned
	//from ParseAll
	NeededBy []string

	HasLitFile bool
	litNode    *lit.Node

	globalResultsMap map[string]*ParseResult
}

func (p *ParseResult) LitHTML() string {
	var b bytes.Buffer
	lit.WriteHTML(&b, p.litNode, &lit.WriteOpts{Prefix: "", Indent: " "})
	return b.String()
}

func (p *ParseResult) MacrosHTML() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "<!-- %d %d -->", len(p.AllNeeds), len(p.NeedsParsed))
	for _, need := range p.AllNeeds {
		pr := p.globalResultsMap[need]
		for _, l := range pr.MacrosLines {
			fmt.Fprintf(&b, "\n\n<!-- %s -->\n\n", l)
			//			fmt.Fprint(&b, "\\(")
			fmt.Fprint(&b, l)
			//			fmt.Fprint(&b, "\\)\n")
		}
	}
	return b.String()
}

func allNeeds(p *ParseResult, all map[string]*ParseResult) []string {
	an := make([]string, 0)
	needs := make([]string, len(p.Needs))
	seen := map[string]bool{}
	copy(needs, p.Needs)
	for len(needs) > 0 {
		n := needs[0]
		needs = needs[1:]
		if seen[n] {
			continue
		}
		seen[n] = true
		an = append(an, n)
		for _, nn := range all[n].Needs {
			needs = append(needs, nn)
		}
	}

	// reverse in place
	// works?
	for i, j := 0, len(an)-1; i < j; i, j = i+1, j-1 {
		an[i], an[j] = an[j], an[i]
	}
	return an
}

type ByName []*ParseResult

func (s ByName) Len() int           { return len(s) }
func (s ByName) Less(i, j int) bool { return s[i].Name < s[j].Name }
func (s ByName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func Parse(sheet io.Reader, macros io.Reader) *ParseResult {
	p := new(ParseResult)

	scanner := bufio.NewScanner(macros)
	for scanner.Scan() {
		t := scanner.Text()
		p.MacrosLines = append(p.MacrosLines, t)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "bbk.Parse: reading macros %v\n", err)
	}

	scanner = bufio.NewScanner(sheet)

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
		fmt.Fprintf(os.Stderr, "bbk.Parse: reading sheet %v\n", err)
	}
	p.Body = strings.Join(p.Lines, " ")
	p.Terms = Terms(p.Body)
	p.AST = ParseSheetAST(p.Lines)
	return p
}

func ParseAll(sheetsdir string) ([]*ParseResult, error) {
	files, err := ioutil.ReadDir(sheetsdir)
	if err != nil {
		return nil, err
	}
	results := make([]*ParseResult, 0, len(files))
	m := make(map[string]*ParseResult)
	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		sheetfile, err := os.Open(filepath.Join(sheetsdir, f.Name(), "sheet.tex"))
		if os.IsNotExist(err) {
			log.Printf("bbk.ParseAll: directory %q lacks sheets.tex", f.Name())
			continue
		} else if err != nil {
			log.Fatalf("bbk.ParseAll: opening sheet.tex %v", err)
		}
		macrosfile, err := os.Open(filepath.Join(sheetsdir, f.Name(), "macros.tex"))
		if os.IsNotExist(err) {
			log.Printf("bbk.ParseAll: directory %q lacks macros.tex", f.Name())
			continue
		} else if err != nil {
			log.Fatalf("bbk.ParseAll: opening macros.tex %v", err)
		}

		p := Parse(sheetfile, macrosfile)
		if p.Name == "" {
			log.Fatalf("no name for file %v", f)
		}
		results = append(results, p)
		m[p.Name] = p
		sheetfile.Close()
		macrosfile.Close()

		lf, err := os.Open(filepath.Join(sheetsdir, f.Name(), "sheet.lit"))
		if err == nil {
			p.HasLitFile = true
			bs, err := ioutil.ReadAll(lf)
			if err != nil {
				log.Fatal(err)
			}
			n, err := lit.ParseLit(string(bs))
			if err != nil {
				log.Fatal(err)
			}
			p.litNode = n
		}
	}

	for _, p := range results {
		for _, n := range p.Needs {
			o, ok := m[n]
			if !ok {
				log.Fatalf("bbk.ParseAll: %q references missing %q", p.Name, n)
			}

			o.NeededBy = append(o.NeededBy, p.Name)
		}
	}

	sort.Sort(ByName(results))

	for _, p := range results {
		sort.Strings(p.Needs)
		sort.Strings(p.NeededBy)
	}
	for _, p := range results {
		p.AllNeeds = allNeeds(p, m)
		p.globalResultsMap = m
	}
	return results, nil
}

func Terms(s string) []string {
	original := s
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
	s = original
	for {
		i := strings.Index(s, "\\t{")
		if i == -1 {
			break
		}
		s = s[i+3:]
		j := strings.Index(s, "}")
		ts = append(ts, s[:j])
	}
	return ts
}

func ParseSheetAST(lines []string) *SheetAST {
	var p parser
	var t SheetAST
	for _, l := range lines {
		l = strings.TrimSpace(l)

		if strings.HasPrefix(l, namePrefix) {
			t.Name = strings.TrimPrefix(l, namePrefix)
			continue
		}

		if strings.HasPrefix(l, "%") {
			continue
		}

		if strings.HasPrefix(l, "\\ssection{") {
			// s has "}"
			s := strings.TrimPrefix(l, "\\ssection{")
			// drop "}"
			s = s[:len(s)-1]
			t.Nodes = append(t.Nodes, SheetSection(s))
			continue
		}
		if strings.HasPrefix(l, "\\ssubsection{") {
			// s has "}"
			s := strings.TrimPrefix(l, "\\ssubsection{")
			// drop "}"
			s = s[:len(s)-1]
			t.Nodes = append(t.Nodes, SheetSubSection(s))
			continue
		}

		if l == "" {
			p.paragraph = nil
			continue
		}

		if p.paragraph == nil {
			p.paragraph = new(SheetParagraph)
			t.Nodes = append(t.Nodes, p.paragraph)
		}

		p.paragraph.Sentences = append(p.paragraph.Sentences, l)
	}
	return &t
}

type parser struct {
	paragraph *SheetParagraph
}

type SheetAST struct {
	Name  string
	Nodes []SheetNode
}

var emphregex = regexp.MustCompile(`\\emph\{(.+)\}`)
var textbfregex = regexp.MustCompile(`\\textbf\{(.+)\}`)
var textitregex = regexp.MustCompile(`\\textit\{(.+)\}`)
var ctregex = regexp.MustCompile(`\\ct\{(.+)\}\{.*\}`)
var sayregex = regexp.MustCompile(`\\say\{(.+)\}`)

func (a *SheetAST) String() string {
	var numSection = 1
	var numSubSection = 1
	var b strings.Builder
	for _, n := range a.Nodes {
		switch v := n.(type) {
		case SheetSection:
			fmt.Fprintf(&b, `<div class="sheet-listing" id="%s"><h1 class="sheet">%d. %s</h1></div>`, DirName(string(v)), numSection, v)
			numSubSection = 1
			numSection += 1
		case SheetSubSection:
			fmt.Fprintf(&b, `<div class="sheet-listing"><h2 class="sheet" id="%s">%d.%d %s</h2></div>`, DirName(string(v)), numSection-1, numSubSection, v)
			numSubSection += 1
		case *SheetParagraph:
			fmt.Fprint(&b, `<div class="sheet-listing"><p class="sheet">`)
			for _, s := range v.Sentences {
				s = emphregex.ReplaceAllString(s, "<i>$1</i>")
				s = textitregex.ReplaceAllString(s, "<i>$1</i>")
				s = textbfregex.ReplaceAllString(s, "<b>$1</b>")
				s = ctregex.ReplaceAllString(s, "<i>$1</i>")
				s = sayregex.ReplaceAllString(s, "\"$1\"")
				fmt.Fprintf(&b, " %s", s)
			}
			fmt.Fprint(&b, `</p></div>`)
		}
	}
	return b.String()
}

type SheetNode interface{ IsSheetNode() bool }

type SheetSection string

func (s SheetSection) IsSheetNode() bool { return true }

type SheetSubSection string

func (s SheetSubSection) IsSheetNode() bool { return true }

type SheetParagraph struct {
	Sentences []string
}

func (p *SheetParagraph) IsSheetNode() bool { return true }

func init() {
	gob.Register(SheetSection(""))
	gob.Register(SheetSubSection(""))
	gob.Register(&SheetParagraph{})
}

type SheetSentence string

func (s SheetSentence) IsSheetNode() bool { return true }
