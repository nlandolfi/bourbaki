package bbk

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/nlandolfi/lit"
	"gopkg.in/yaml.v3"
)

func (s *Sheet) LitHTML() string {
	var b bytes.Buffer
	lit.WriteHTML(&b, s.LitNode, &lit.WriteOpts{Prefix: "", Indent: " "})
	return b.String()
}

func AllNeeds(sheet *Sheet, ss *SheetSet) []string {
	an := make([]string, 0)
	needs := make([]string, len(sheet.Config.Needs))
	seen := map[string]bool{}
	copy(needs, sheet.Config.Needs)
	for len(needs) > 0 {
		n := needs[0]
		needs = needs[1:]
		if seen[n] {
			continue
		}
		seen[n] = true
		an = append(an, n)
		for _, nn := range ss.Sheets[n].Config.Needs {
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

func MacrosOrEmpty(sheet *Sheet) string {
	for c := sheet.LitNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == lit.CommentNode && len(c.Data) > 5 && c.Data[:6] == "macros" {
			return strings.TrimSpace(c.Data[strings.Index(c.Data, "\n"):])
		}
	}
	return ""
}

type SheetConfig struct {
	Name      string   `yaml:"name",omitempty`
	Needs     []string `yaml:"needs,omitempty"`
	Refs      []string `yaml:"refs,omitempty"`
	Wikipedia string   `yaml:"wikipedia,omitempty"`
}

type Sheet struct {
	Config SheetConfig
	// this lit node has its yaml sripped
	LitNode *lit.Node
}

type SheetDir struct {
	Sheet         *Sheet
	MacrosTexFile string
}

func ParseSheetDir(dir fs.FS) (*SheetDir, error) {
	var sd SheetDir
	sf, err := dir.Open("sheet.lit")
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
	sd.Sheet = s

	mf, err := dir.Open("macros.tex")
	if err != nil {
		return nil, err
	}
	bs, err := io.ReadAll(mf)
	if err != nil {
		return nil, err
	}
	sd.MacrosTexFile = string(bs)
	if err := mf.Close(); err != nil {
		return nil, err
	}
	return &sd, nil
}

func WriteSheet(w io.Writer, s *Sheet) error {
	header, err := yaml.Marshal(s.Config)
	if err != nil {
		return err
	}

	// hack TODO fix - NCL 2/25/2023; brittle to how lit does it
	if _, err := fmt.Fprintf(w, "<!--yaml\n%s-->\n\n", header); err != nil {
		return err
	}

	if err := lit.WriteLit(w, s.LitNode, lit.DefaultWriteOpts); err != nil {
		return fmt.Errorf("fmt.Fprintf: %w", err)
	}

	return nil
}

func OverwriteSheetDir(dir string, sd *SheetDir) error {
	// ensure the dir exists
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("os.Mkdir error: %w", err)
		}
	} // handle error?

	// write the Makefile
	tmpl := template.Must(template.New("").Parse(MakefileTemplate))
	f, err := os.Create(path.Join(dir, "Makefile"))
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	if err := tmpl.Execute(f, sd.Sheet); err != nil {
		return fmt.Errorf("tmpl.Execute: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("file.Close: %w", err)
	}

	// write the macros.tex file
	f, err = os.Create(path.Join(dir, "macros.tex"))
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	if _, err := f.Write([]byte(sd.MacrosTexFile)); err != nil {
		return fmt.Errorf("fmt.Fprintf: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("file.Close: %w", err)
	}

	// write the sheet lit file
	f, err = os.Create(path.Join(dir, "sheet.lit"))
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	if err := WriteSheet(f, sd.Sheet); err != nil {
		return fmt.Errorf("WriteSheet: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("file.Close: %w", err)
	}

	return nil
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

	// easier to think about not have the data in two places
	n.RemoveChild(n.FirstChild)

	return s, nil
}

var termsRegex = regexp.MustCompile("❬([^❭]*?)❭")

func ParseTerms(s *Sheet) ([]string, error) {
	// hack for now
	var b bytes.Buffer
	if err := lit.WriteLit(&b, s.LitNode, &lit.WriteOpts{Prefix: "", Indent: ""}); err != nil {
		return nil, err
	}
	matches := termsRegex.FindAllStringSubmatch(b.String(), -1) // Find all matches in text

	var out []string = make([]string, 0, len(matches))

	for _, match := range matches {
		// nonsense for hack, later this should all use lit's tokenization
		unclean := strings.Replace(match[1], "\n", " ", -1)

		pieces := strings.Split(unclean, " ")
		for i, p := range pieces {
			pieces[i] = strings.TrimSpace(p)
		}

		out = append(out, strings.Join(pieces, " ")) //  the matched text (capturing group 1)
	}
	return out, nil
}

type SheetSet struct {
	Sheets map[string]*Sheet
}

type SheetsDir struct {
	Sheets map[string]*SheetDir
}

// Use ParseSheets to parse all the sheet directories in dir.
func ParseSheets(dir fs.FS) (*SheetsDir, error) {
	var ss SheetsDir
	ss.Sheets = make(map[string]*SheetDir)

	entries, err := fs.ReadDir(dir, ".")
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		sdfs, err := fs.Sub(dir, e.Name())
		if err != nil {
			return nil, err
		}
		sd, err := ParseSheetDir(sdfs)
		if err != nil {
			return nil, err
		}
		if sd.Sheet.Config.Name != e.Name() {
			return nil, fmt.Errorf("sheet name: %q doesn't match dir name: %q", sd.Sheet.Config.Name, e.Name())

		}
		if _, ok := ss.Sheets[sd.Sheet.Config.Name]; ok {
			return nil, fmt.Errorf("duplicate sheet name: %q on dir: %q", sd.Sheet.Config.Name, e.Name())
		}

		ss.Sheets[sd.Sheet.Config.Name] = sd
	}
	return &ss, nil
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
