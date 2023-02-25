package bbk

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/nlandolfi/lit"
	"gopkg.in/yaml.v3"
)

type ParseResult struct {
	Lines       []string
	Body        string
	NeedsParsed []*ParseResult
	Terms       []string
	MacrosLines []string
	AllNeeds    []string

	// only set if returned
	//from ParseAll
	NeededBy []string

	HasLitFile bool
	litNode    *lit.Node
	Config     SheetConfig

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
	needs := make([]string, len(p.Config.Needs))
	seen := map[string]bool{}
	copy(needs, p.Config.Needs)
	for len(needs) > 0 {
		n := needs[0]
		needs = needs[1:]
		if seen[n] {
			continue
		}
		seen[n] = true
		an = append(an, n)
		for _, nn := range all[n].Config.Needs {
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
func (s ByName) Less(i, j int) bool { return s[i].Config.Name < s[j].Config.Name }
func (s ByName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func Parse(sheet, macros, lf io.Reader) *ParseResult {
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
		p.Lines = append(p.Lines, t)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "bbk.Parse: reading sheet %v\n", err)
	}
	p.Body = strings.Join(p.Lines, " ")
	p.Terms = Terms(p.Body)
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
	if p.litNode.FirstChild.Type == lit.YAMLNode {
		if err := yaml.Unmarshal([]byte(n.FirstChild.Data), &p.Config); err != nil {
			log.Fatal(err)
		}
	}
	return p
}

type SheetConfig struct {
	Name      string   `yaml:"name",omitempty`
	Needs     []string `yaml:"needs,omitempty"`
	Refs      []string `yaml:"refs,omitempty"`
	Wikipedia string   `yaml:"wikipedia,omitempty"`
}

type Sheet struct {
	Config  SheetConfig
	LitNode *lit.Node
}

// Use ParseSheet to parse a sheet.lit file
func ParseSheet(r io.Reader) (*Sheet, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	n, err := lit.ParseLit(string(bs))
	if err != nil {
		return nil, err
	}
	s := &Sheet{
		LitNode: n,
	}

	if n.FirstChild.Type != lit.YAMLNode {
		return nil, fmt.Errorf("first node should be yaml")
	}
	if err := yaml.Unmarshal([]byte(n.FirstChild.Data), &s.Config); err != nil {
		return nil, err
	}

	return s, nil
}

type SheetSet struct {
	Sheets map[string]*Sheet
}

// Use ParseSheetSet to parse all the sheet directories in dir.
func ParseSheetSet(dir fs.FS) (*SheetSet, error) {
	var ss SheetSet
	ss.Sheets = make(map[string]*Sheet)

	entries, err := fs.ReadDir(dir, ".")
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		sf, err := dir.Open(path.Join(e.Name(), "sheet.lit"))
		if err != nil {
			return nil, err
		}
		s, err := ParseSheet(sf)
		if err != nil {
			return nil, err
		}
		if err := sf.Close(); err != nil {
			return nil, err
		}
		if s.Config.Name != e.Name() {
			return nil, fmt.Errorf("sheet name: %q doesn't match dir name: %q", s.Config.Name, e.Name())

		}
		if _, ok := ss.Sheets[s.Config.Name]; ok {
			return nil, fmt.Errorf("duplicate sheet name: %q on dir: %q", s.Config.Name, e.Name())
		}
		ss.Sheets[s.Config.Name] = s
	}
	return &ss, nil
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
		litfile, err := os.Open(filepath.Join(sheetsdir, f.Name(), "sheet.lit"))
		if os.IsNotExist(err) {
			log.Printf("bbk.ParseAll: directory %q lacks sheet.lit", f.Name())
			continue
		} else if err != nil {
			log.Fatalf("bbk.ParseAll: opening sheet.lit %v", err)
		}

		p := Parse(sheetfile, macrosfile, litfile)
		if err := sheetfile.Close(); err != nil {
			log.Fatal(err)
		}
		if err := macrosfile.Close(); err != nil {
			log.Fatal(err)
		}
		if err := litfile.Close(); err != nil {
			log.Fatal(err)
		}

		if p.Config.Name == "" {
			log.Fatalf("no name for file %v", f)
		}
		m[p.Config.Name] = p
		results = append(results, p)

	}

	for _, p := range results {
		for _, n := range p.Config.Needs {
			o, ok := m[n]
			if !ok {
				log.Fatalf("bbk.ParseAll: %q references missing %q", p.Config.Name, n)
			}

			o.NeededBy = append(o.NeededBy, p.Config.Name)
		}
	}

	sort.Sort(ByName(results))

	for _, p := range results {
		sort.Strings(p.Config.Needs)
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
