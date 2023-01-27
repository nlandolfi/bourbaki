package main

import (
	"bbk"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

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
	default:
		s.info("no detailed help for %q", v)
	}
}
func checkMain(s *state) {
	// for now must run from sheets directory
	ss, err := bbk.ParseSheetSet(os.DirFS("."))
	if err != nil {
		log.Fatal(err)
	}
	//	log.Printf("%+v", ss.Sheets)
	log.Printf("%d sheets", len(ss.Sheets))
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
