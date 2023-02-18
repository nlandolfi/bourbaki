package main

import (
	"bbk"
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/nlandolfi/lit"
	"gopkg.in/yaml.v3"
)

// Set using link flags; e.g., -X main.Version=...
var (
	Version   string // e.g. 0.1.0
	GitSHA    string
	BuildDate string
)

// the help string that is printed for command `bbk`
const basicHelp = `bbk <command>
    - check
    - terms 
    - help <command>
    - version`

func main() {
	s := &state{
		Args: os.Args,
		In:   bufio.NewReader(os.Stdin),
		Out:  os.Stdout,
		Err:  os.Stderr,
	}
	if len(s.Args) < 2 {
		s.info(basicHelp)
		return
	}
	switch cmd := s.Args[1]; cmd {
	case "help", "helps":
		helpMain(s)
	case "check":
		checkMain(s)
	case "test", "t":
		testMain(s)
	case "terms":
		termsMain(s)
	case "graph":
		graphMain(s)
	case "mv":
		mvMain(s)
	case "sheets":
		sheetsMain(s)
	case "version", "v":
		s.info("bbk version %s \n  SHA %s \n  Built at %s", Version, GitSHA, BuildDate)
	default:
		fmt.Printf("unknown subcommand: %q", cmd)
		fmt.Printf("\n")
	}
}

// state is a helper type for the information
// needed to process a command; it has helper methods
type state struct {
	Args []string
	In   *bufio.Reader
	Err  io.Writer
	Out  io.Writer
	Config
}

// print and exit
func (s *state) error(f string, vs ...interface{}) {
	fmt.Fprintf(s.Err, "ERROR:"+f+"\n", vs...)
	os.Exit(1)
}

// print some info
func (s *state) info(f string, vs ...interface{}) {
	fmt.Fprintf(s.Out, f+"\n", vs...)
}

// get some input, after printing a formatted string
func (s *state) input(f string, vs ...interface{}) string {
	fmt.Fprintf(s.Out, f, vs...)
	out, err := s.In.ReadString('\n')
	if err != nil {
		s.error("reading input: %v", err)
	}
	return strings.TrimSuffix(out, "\n")
}

type Config struct {
	Public  string
	Private string
}

// helpMain handles printing a help menu
func helpMain(s *state) {
	if len(s.Args) < 3 {
		s.info("bbk help <command>")
		return
	}
	switch v := s.Args[2]; v {
	case "terms":
		s.info(termsHelp)
	case "mv":
		s.info(mvHelp)
	case "check":
		s.info(checkHelp)
	default:
		s.info("no detailed help for %q", v)
	}
}

const checkHelp = `bbk check

Parse all sheets. Run from sheets directory.

Examples
  - bbk check
`

func checkMain(s *state) {
	// for now must run from sheets directory
	st := time.Now()
	ss, err := bbk.ParseSheetSet(os.DirFS("."))
	if err != nil {
		log.Fatal(err)
	}
	//	log.Printf("%+v", ss.Sheets)
	s.info("%d sheets parsed in %d", len(ss.Sheets), time.Now().Sub(st))
}

func testMain(s *state) {
	// for now must run from sheets directory
	ss, err := bbk.ParseSheetSet(os.DirFS("."))
	if err != nil {
		log.Fatal(err)
	}
	//	log.Printf("%+v", ss.Sheets)
	log.Printf("%d sheets", len(ss.Sheets))

	for name, s := range ss.Sheets {
		if s.ConfigFromComment != nil {
			bs, err := yaml.Marshal(s.Config)
			if err != nil {
				log.Fatal(err)
			}
			y := make(map[interface{}]interface{})
			if err := yaml.Unmarshal(bs, &y); err != nil {
				log.Fatal(err)
			}
			s.LitNode.InsertBefore(&lit.Node{
				Type: lit.YAMLNode,
				// wont have data but who cares
				YAML:      y,
				IsComment: true,
			}, s.ConfigFromComment)
			s.LitNode.RemoveChild(s.ConfigFromComment)
		}
		f, err := os.Create(path.Join("./", name, "sheet.lit"))
		if err != nil {
			log.Fatal(err)
		}
		// TODO: shouldn't WriteLit possibly return errors?? -NCL 1/25/23
		lit.WriteLit(f, s.LitNode, lit.DefaultWriteOpts)
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

const termsHelp = `bbk terms
`

func termsMain(s *state) {
	log.Print("not implemented")
}

const graphHelp = `bbk terms
`

func graphMain(s *state) {
	log.Print("not implemented")
}

const mvHelp = `bbk mv <from> <to>

Change the name of a sheet. Useful because it will find and 
replace the name in all other sheets as well.

Examples
  - bbk mv markov_chains memory_chains
`

func mvMain(s *state) {
	if len(s.Args) < 4 {
		s.info(mvHelp[:strings.Index(mvHelp, "\n")])
		return
	}

	from, to := s.Args[2], s.Args[3]

	_, err := bbk.ParseSheetSet(os.DirFS("."))
	if err != nil {
		s.error("parsing sheet set: %v", err)
	}

	log.Printf("not implemented: %s -> %s", from, to)
}

const sheetsHelp = `bbk sheets

This is the old bbk_sheets command.

Examples
  - bbk sheets
`

func sheetsMain(s *state) {
	fs := flag.NewFlagSet("sheets", flag.ContinueOnError)

	var (
		sheetsDir = fs.String("sheets", ".", "the sheets directory")
		startAt   = fs.String("start-at", "", "sheet to start compiling at (alphabetical)")
		dry       = fs.Bool("dry", false, "whether to actually run") // use just to check lit status
	)

	if err := fs.Parse(s.Args[2:]); err != nil {
		s.error("parsing flags: %v", err)
		return
	}

	rs, err := bbk.ParseAll(*sheetsDir)

	if err != nil {
		log.Fatalf("bbk.ParseAll: %v", err)
	}
	results := make(map[string]*bbk.ParseResult, len(rs))
	for _, p := range rs {
		results[p.Config.Name] = p
	}

	if len(results) == 0 {
		log.Printf("warning: no sheets parsed")
	}

	names := make([]string, 0, len(results))
	for n := range results {
		names = append(names, n)
	}
	sort.Strings(names)

	var missingLit int
	for _, name := range names {
		r := results[name]
		if !r.HasLitFile {
			log.Print(r.Config.Name)
			missingLit += 1
		}
	}
	log.Printf("%d/%d missing lit files", missingLit, len(results))

	if *dry {
		return
	}

	/*
		for _, name := range names {
			f, err := os.Open(path.Join(*sheetsDir, name, "sheet.tex"))
			if err != nil {
				log.Fatal(err)
			}
			h := sha256.New()
			if _, err := io.Copy(h, f); err != nil {
				log.Fatal(err)
			}
			f.Close()

			fmt.Printf("%x", h.Sum(nil))

			currentHashFile, err := os.Open(
			f, err = os.Create(fmt.Sprintf("/tmp/bbk_sheets_%s",
				name))

		}
	*/

	tmpl := template.Must(template.New("").Parse(bbk.MakefileTemplate))
	for _, name := range names {
		out, err := os.Create(path.Join(*sheetsDir, name, "Makefile"))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		if err := tmpl.Execute(out, results[name]); err != nil {
			log.Fatal(err)
		}
		if err := out.Close(); err != nil {
			log.Fatal(err)
		}
	}

	/*
		for name := range results {
			c := exec.Command("make", "remake")
			c.Dir = name
			bs, err := c.CombinedOutput()
			if err != nil {
				fmt.Print(string(bs))
				log.Fatalf("%q: %v", name, err)
			}
			c = exec.Command("make")
			c.Dir = name
			bs, err = c.CombinedOutput()
			if err != nil {
				fmt.Print(string(bs))
				log.Fatalf("%q: %v", name, err)
			}
		}
	*/

	var wg sync.WaitGroup
	ch := make(chan string, len(results))

	// these are the workers
	for i := 0; i < 32; i++ {
		go func(in <-chan string) {
			for name := range ch {
				if *startAt > name {
					wg.Done()
					continue
				}
				c := exec.Command("make", "remake")
				c.Dir = name
				bs, err := c.CombinedOutput()
				if err != nil {
					fmt.Print(string(bs))
					log.Fatalf("%q: %v", name, err)
				}
				c = exec.Command("make")
				c.Dir = name
				bs, err = c.CombinedOutput()
				if err != nil {
					fmt.Print(string(bs))
					log.Fatalf("%q: %v", name, err)
				}
				log.Printf("done: %q", name)
				wg.Done()
			}
		}(ch)
	}

	for _, n := range names {
		wg.Add(1)
		ch <- n
	}

	close(ch)

	wg.Wait()
}
