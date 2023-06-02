package main

import (
	"bbk"
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/nlandolfi/lit"
	"github.com/nlandolfi/spin/infra/audit"
)

// Set using link flags; e.g., -X main.Version=...
var (
	Version   string // e.g. 0.1.0
	GitSHA    string
	BuildDate string
	GoVersion = runtime.Version()
)

// the help string that is printed for command `bbk`
const basicHelp = `bbk <command>
    - check
    - mk <name>
    - mv <from> <to>
    - rm <from> <to>
    - terms 
    - sheets
    - graph
    - all
    - macros
		- compile_screen
		- serve
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
	case "terms":
		termsMain(s)
	case "graph":
		graphMain(s)
	case "rm":
		rmMain(s)
	case "mv":
		mvMain(s)
	case "mk":
		mkMain(s)
	case "sheets":
		sheetsMain(s)
	case "sheet":
		sheetMain(s)
	case "macros":
		macrosMain(s)
	case "compile_screen":
		compileScreenMain(s)
	case "serve":
		serveMain(s)
	case "all":
		allMain(s)
	case "rand":
		randMain(s)
	case "version", "v":
		s.info("bbk version %s (%s)\n  SHA %s \n  Built at %s", Version, GoVersion, GitSHA, BuildDate)
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
	case "rm":
		s.info(rmHelp)
	case "mk":
		s.info(mkHelp)
	case "check":
		s.info(checkHelp)
	case "sheets":
		s.info(sheetsHelp)
	case "all":
		s.info(allHelp)
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
	s.info("%d sheets parsed in %s", len(ss.Sheets), time.Now().Sub(st).String())
}

func serveMain(s *state) {
	fs := flag.NewFlagSet("sheet", flag.ContinueOnError)

	var (
		staticDir  = fs.String("static-dir", "/static", "directory of static files")
		localAudit = fs.Bool("local-audit", false, "only locally log audit")
		auditURL   = fs.String("audit-url", "https://audit-5rcmpbrboq-uw.a.run.app", "audit service to use")
	)

	/* TODO: bring back audit log
	var noKey = &spin.Key{
		Name: "bbk-no-key",
		Type: "bbk-no-key",
	}
	*/

	if err := fs.Parse(os.Args[2:]); err != nil {
		log.Fatal(err)
	}

	var auditLogger audit.Logger
	if *localAudit {
		auditLogger = &audit.LocalLogger{log.Default()}
	} else {
		auditLogger = audit.NewClient(*auditURL)
	}
	as := &audit.Service{
		Name:   "bbk-serve",
		Logger: auditLogger,
	}

	f, err := os.Open(path.Join(*staticDir, "results.json.gzip"))
	if err != nil {
		log.Fatal(err)
	}
	r, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	var sd bbk.SearchData
	if err := json.NewDecoder(r).Decode(&sd); err != nil {
		log.Fatal(err)
	}
	var results = make(map[string]*bbk.SheetResult, len(sd.Results))
	for _, r := range sd.Results {
		results[r.Config.Name] = r
	}
	srcher := &searcher{sd.Version, bbk.NewSearcher(results, *staticDir), as}

	m := http.NewServeMux()
	m.HandleFunc("/search", srcher.search)
	m.HandleFunc("/search_json", srcher.search_json)

	fsfs := http.FileServer(http.Dir(*staticDir))
	m.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fsfs.ServeHTTP(w, r)

		// LOGGING
		data := map[string]string{
			"remote-addr": r.RemoteAddr,
			"url":         r.URL.String(),
			"version":     sd.Version,
		}
		for k, vs := range r.Header {
			data[k] = strings.Join(vs, ",,")
		}
		// TODO: bring back audit log
		//srcher.AuditLog.Put("get", noKey, data)
	}))
	log.Fatal(http.ListenAndServe(":8080", m))
}

type searcher struct {
	Version string
	*bbk.Searcher
	AuditLog *audit.Service
}

func (s *searcher) data(r *http.Request) *SearchData {
	start := time.Now()
	var sd SearchData
	sd.Version = s.Version
	sd.RemoteAddr = r.RemoteAddr
	sd.Query = r.FormValue("query")
	sd.RewrittenQuery = bbk.RewriteQuery(sd.Query)
	if sd.RewrittenQuery != "" {
		sd.Results = s.Search(sd.RewrittenQuery)
		sd.Searched = true
	}
	sd.SearchDuration = time.Now().Sub(start)
	return &sd
}

func (s *searcher) log(action string, sd *SearchData, r *http.Request) {
	data := map[string]string{
		"remote-addr":     sd.RemoteAddr,
		"query":           sd.Query,
		"rewritten-query": sd.RewrittenQuery,
		"duration":        fmt.Sprintf("%d", sd.SearchDuration),
		"version":         sd.Version,
	}
	for k, vs := range r.Header {
		data[k] = strings.Join(vs, ",,")
	}
	// TODO: bring back audit log
	//s.AuditLog.Put("search", noKey, data)
}

func (s *searcher) search(w http.ResponseWriter, r *http.Request) {
	sd := s.data(r)
	searchTemplate := template.Must(
		template.New("search.tmpl").Funcs(
			template.FuncMap{
				"title":   bbk.Title,
				"reasons": reasons,
			},
		).ParseFiles(path.Join(s.Searcher.StaticDir, "search.tmpl")))
	if err := searchTemplate.Execute(w, sd); err != nil {
		http.Error(w, "internal", 500)
		log.Fatal(err)
	}
	s.log("search", sd, r)
}

func (s *searcher) search_json(w http.ResponseWriter, r *http.Request) {
	sd := s.data(r)
	if err := json.NewEncoder(w).Encode(sd); err != nil {
		http.Error(w, "internal", 500)
		log.Fatal(err)
	}
	s.log("search-json", sd, r)
}

type SearchData struct {
	Version        string
	RemoteAddr     string
	Query          string
	RewrittenQuery string
	Results        []*bbk.SearchResult
	Searched       bool
	SearchDuration time.Duration
}

func reasons(sr *bbk.SearchResult) string {
	return strings.Join(sr.Reasons, ", ")
}

func compileScreenMain(s *state) {
	fs := flag.NewFlagSet("sheet", flag.ContinueOnError)

	sheetsDir := fs.String("sheets-dir", "../sheets", "the sheets directory")
	staticDir := fs.String("static-dir", "./static", "directory of static files")

	if err := fs.Parse(os.Args[2:]); err != nil {
		log.Fatal(err)
	}

	ss, err := bbk.ParseSheetSet(os.DirFS(*sheetsDir))
	if err != nil {
		log.Fatal(err)
	}

	if len(ss.Sheets) == 0 {
		log.Printf("warning: no sheets found")
	}

	cmd := exec.Command("git", "rev-parse", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	gitCommit := out.String()

	// log.Printf("%d sheets", len(results))

	sheetTemplate := template.Must(
		template.New("sheet.tmpl").Funcs(
			template.FuncMap{
				"title": bbk.Title,
				"needed_by": func(s *bbk.Sheet) (out []string) {
					for _, os := range ss.Sheets {
						for _, n := range os.Config.Needs {
							if n == s.Config.Name {
								out = append(out, os.Config.Name)
							}
						}
					}
					sort.Strings(out)
					return
				},
			},
		).ParseFiles(path.Join(*staticDir, "sheet.tmpl")))

	type SheetData struct {
		*bbk.Sheet
		Version string
	}

	for _, p := range ss.Sheets {
		name := p.Config.Name
		out, err := os.Create(filepath.Join(*staticDir, "sheets", fmt.Sprintf("%s.html", name)))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		if err := sheetTemplate.Execute(out, &SheetData{p, gitCommit[:9]}); err != nil {
			log.Fatal(err)
		}
		out.Close()

		_, err = bbk.CopyFile(
			filepath.Join(*sheetsDir, name, fmt.Sprintf("%s.pdf", name)),
			filepath.Join(*staticDir, "sheets", fmt.Sprintf("%s.pdf", name)),
		)
		if err != nil {
			log.Printf("error transferring %s/sheet.pdf: %v", name, err)
		}

		_, err = bbk.CopyFile(
			filepath.Join(*sheetsDir, name, "graph.pdf"),
			filepath.Join(*staticDir, "sheets", fmt.Sprintf("%s-graph.pdf", name)),
		)
		if err != nil {
			log.Printf("error transferring %s/graph.pdf: %v", name, err)
		}

		// BUG: if two of the same named graphics, then an error
		graphicsDir := filepath.Join(*sheetsDir, name, "graphics")
		_, err = os.Stat(graphicsDir)
		var hasGraphics bool
		if os.IsNotExist(err) {
		} else if err != nil {
		} else {
			hasGraphics = true
		}
		if hasGraphics {
			toDir := filepath.Join(*staticDir, "sheets", "graphics")
			os.Mkdir(toDir, 0700)
			fs, err := ioutil.ReadDir(graphicsDir)
			if err != nil {
				log.Fatal(err)
			}
			for _, fi := range fs {
				_, err = bbk.CopyFile(
					filepath.Join(graphicsDir, fi.Name()),
					filepath.Join(toDir, fi.Name()),
				)
				if err != nil {
					log.Printf("error transferring %s/%s: %v", p, fi.Name(), err)
				}
			}
		}
	}

	type IndexData struct {
		Results []*bbk.Sheet
		Version string
	}

	indexTemplate := template.Must(
		template.New("index.tmpl").Funcs(
			template.FuncMap{
				"title": bbk.Title,
			},
		).ParseFiles(filepath.Join(*staticDir, "index.tmpl")))

	f, err := os.Create(filepath.Join(*staticDir, "index.html"))
	if err != nil {
		log.Fatalf("os.Create index.html: %v", err)
	}
	var results []*bbk.Sheet = make([]*bbk.Sheet, 0, len(ss.Sheets))
	for _, s := range ss.Sheets {
		results = append(results, s)
	}
	sort.Sort(SheetsByName(results))
	if err := indexTemplate.Execute(f, &IndexData{results, gitCommit[:9]}); err != nil {
		log.Fatalf("executing index template: %v", err)
	}
	f.Close()

	f, err = os.Create(filepath.Join(*staticDir, "results.json.gzip"))
	if err != nil {
		log.Fatalf("os.Create results.json.gzip: %v", err)
	}
	var rs []*bbk.SheetResult = make([]*bbk.SheetResult, len(results))
	var b bytes.Buffer
	for i, r := range results {
		b.Reset()
		lit.WriteTex(&b, r.LitNode, &lit.WriteOpts{Prefix: "", Indent: " "})
		ts, err := bbk.ParseTerms(r)
		if err != nil {
			log.Fatal(err)
		}
		rs[i] = &bbk.SheetResult{
			Config: r.Config,
			Terms:  ts,
			RawTex: b.String(),
		}
	}
	w := gzip.NewWriter(f)
	if err := json.NewEncoder(w).Encode(&bbk.SearchData{rs, gitCommit[:9]}); err != nil {
		log.Fatalf("json encoding results: %v", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("gzip writer close: %v", err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("results.json.gzip close: %v", err)
	}
}

type SheetsByName []*bbk.Sheet

func (a SheetsByName) Len() int           { return len(a) }
func (a SheetsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SheetsByName) Less(i, j int) bool { return a[i].Config.Name < a[j].Config.Name }

const termsHelp = `bbk terms
`

func termsMain(s *state) {
	log.Print("not implemented")
}

const graphHelp = `bbk graph <-csv graph.csv> <-tmpl graph.tmpl>

This is the old bbk_graph command.
`

func graphMain(s *state) {
	fs := flag.NewFlagSet("graph", flag.ContinueOnError)
	log.Print("not implemented")
	var (
		inputFileFlag    = fs.String("csv", "graph.csv", "csv file")
		templateFileFlag = fs.String("tmpl", "graph.tmpl", "which template")
	)
	if err := fs.Parse(s.Args[2:]); err != nil {
		s.error("parsing arguments: %v", err)
		return
	}

	tmpl := template.Must(
		template.ParseFiles(*templateFileFlag),
	)

	f, err := os.Open(*inputFileFlag)
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}
	defer f.Close()

	entries, err := bbk.ParseGraph(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.Execute(os.Stdout, entries); err != nil {
		log.Fatal(err)
	}
}

const rmHelp = `bbk rm <from> <to> {-dry} {-sheets <path>}

Like bbk mv, except doesn't overwrite the changed directory.
In other words, this EXPECTS to be a sheet.

WARNING: WILL NOT UPDATE LINKS (as of now 2/25/23)

If you want to remove a directory that has no needs, just

    rm -rf <dir>

Examples
  - bbk rm probability_distributions outcome_probabilities

    All sheets will now point to outcome_probabilties that 
    used to point to probability_distributions, but the
    outcome_probabilities sheet will remain unchanged.
`

func rmMain(s *state) {
	mv_rm_Main(s, false)
}

const mvHelp = `bbk mv <from> <to> {-dry} {-sheets <path>}

Change the name of a sheet. Useful because it will find and 
replace the name in all other sheets as well.

WARNING: WILL NOT UPDATE LINKS (as of now 2/25/23)

Examples
  - bbk mv markov_chains memory_chains
`

func mvMain(s *state) {
	mv_rm_Main(s, true)
}

func mv_rm_Main(s *state, overwriteTo bool) {
	fs := flag.NewFlagSet("mv", flag.ContinueOnError)

	var (
		sheetsDir = fs.String("sheets", ".", "the sheets directory")
		dry       = fs.Bool("dry", false, "whether to actually run")
	)

	if len(s.Args) < 4 {
		s.info(mvHelp[:strings.Index(mvHelp, "\n")])
		return
	}

	from, to := s.Args[2], s.Args[3]

	if err := fs.Parse(s.Args[4:]); err != nil {
		s.error("parsing flags: %v", err)
		return
	}

	ss, err := bbk.ParseSheets(os.DirFS(*sheetsDir))
	if err != nil {
		s.error("parsing sheet set: %v", err)
		return
	}

	fromSheet, ok := ss.Sheets[from]
	if !ok {
		s.error("sheet %s not found", from)
		return
	}
	if overwriteTo {
		if _, ok := ss.Sheets[to]; ok {
			s.error("sheet %s already exists; will not overwrite", to)
			return
		}
	} else {
		if _, ok := ss.Sheets[to]; !ok {
			s.error("sheet %s does not exists; will not perform rm", to)
			return
		}
	}

	// ok, the gameplan is, loop through all sheets
	//    - update needs

	touched := map[string]bool{}

	//var buf bytes.Buffer
	for name, sd := range ss.Sheets {
		if name == from {
			continue
		}

		for i, need := range sd.Sheet.Config.Needs {
			if need == from {
				sd.Sheet.Config.Needs[i] = to
				touched[name] = true
			}
		}

		/*TODO: handle links and other references -- NCL 2/25/23
		buf.Reset()

		// shouldn't this possibly return an error??
		if err := lit.WriteLit(&b1, s.LitNode, lit.DefaultWriteOpts); err != nil {
			c.error("%v", err)
			return
		}

		s := buf.String()
		if strings.Index(s, from) < 0 {
			continue
		}
		out = strings.ReplaceAll(s, from, to)
		nNode, err := lit.ParseLit(out)
		if err != nil {
			c.error("parsing replaced lit: %v", err)
			return
		}
		touched[name] = true
		*/
	}

	if *dry {
		var b bytes.Buffer
		fmt.Fprintf(&b, "would touch:\n")
		tn := make([]string, 0, len(touched))
		for n := range touched {
			tn = append(tn, n)
		}
		sort.Strings(tn)
		for _, n := range tn {
			fmt.Fprintf(&b, " - %s\n", n)
		}
		fmt.Fprintf(&b, "to move %s -> %s", from, to)
		s.info(b.String())
		return
	}
	//s.error("not implemented: %s -> %s", from, to)
	//return

	fromSheet.Sheet.Config.Name = to

	if err := os.RemoveAll(from); err != nil {
		s.error("os.RemoveAll: %v", err)
		return
	}
	for n := range touched {
		if err := bbk.OverwriteSheetDir(n, ss.Sheets[n]); err != nil {
			s.error("os.OverwriteSheetDir: %v", err)
			return
		}
	}
	if overwriteTo {
		if err := bbk.OverwriteSheetDir(to, fromSheet); err != nil {
			s.error("os.OverwriteSheetDir: %v", err)
			return
		}
		c := exec.Command("make")
		c.Dir = to
		bs, err := c.CombinedOutput()
		if err != nil {
			fmt.Print(string(bs))
			s.error("%q: %v", to, err)
			return
		}
	}
}

const mkHelp = `bbk mk <name> {-sheets <path>}

Create the name of a sheet. Useful because it will find and 
replace the name in all other sheets as well.


Examples
  - bbk mk cohomologies
`

func mkMain(s *state) {
	fs := flag.NewFlagSet("mv", flag.ContinueOnError)

	var (
		sheetsDir = fs.String("sheets", ".", "the sheets directory")
	)

	if len(s.Args) < 3 {
		s.info(mkHelp[:strings.Index(mvHelp, "\n")])
		return
	}

	name := s.Args[2]

	if err := fs.Parse(s.Args[3:]); err != nil {
		s.error("parsing flags: %v", err)
		return
	}

	ss, err := bbk.ParseSheets(os.DirFS(*sheetsDir))
	if err != nil {
		s.error("parsing sheet set: %v", err)
		return
	}
	if _, ok := ss.Sheets[name]; ok {
		s.error("%s already exists; see bbk {mv|rm}")
		return
	}

	sd := &bbk.SheetDir{
		Sheet: &bbk.Sheet{
			Config: bbk.SheetConfig{
				Name: name,
			},
			LitNode: blankLitNode,
		},
	}

	if err := bbk.OverwriteSheetDir(name, sd); err != nil {
		s.error("error: %v", err)
		return
	}

	c := exec.Command("make")
	c.Dir = name
	bs, err := c.CombinedOutput()
	if err != nil {
		fmt.Print(string(bs))
		s.error("%q: %v", name, err)
		return
	}
}

var blankLitNode *lit.Node = lit.Must(lit.ParseLit(`
<tex> 
	‖ \blankpage ⦉ 
</tex>
`))

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

	ss, err := bbk.ParseSheetSet(os.DirFS(*sheetsDir))
	if err != nil {
		log.Fatal(err)
	}

	names := make([]string, 0, len(ss.Sheets))
	for n := range ss.Sheets {
		names = append(names, n)
	}
	sort.Strings(names)

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
		if err := tmpl.Execute(out, ss.Sheets[name]); err != nil {
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
	ch := make(chan string, len(ss.Sheets))

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

const sheetHelp = `bbk sheet

This is the old bbk sheet command.

Examples
  - bbk sheet
`

func sheetMain(s *state) {
	fs := flag.NewFlagSet("sheet", flag.ContinueOnError)

	var (
		_ = fs.String("in", "sheet.tex", "sheet file")
		// create, makefile, graph
		mode      = fs.String("mode", "mk", "which mode {c, mk, g, ts}")
		sheetsDir = fs.String("sheets", "../", "where to find other sheets")
	)

	fs.Parse(s.Args[2:])

	ss, err := bbk.ParseSheetSet(os.DirFS(*sheetsDir))
	if err != nil {
		s.error("parsing sheet set: %v", err)
		return
	}

	// just want to get the name
	f, err := os.Open("sheet.lit")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	sheet, err := bbk.ParseSheet(f)
	if err != nil {
		log.Fatal(err)
	}

	switch *mode {
	case "c":
		sheetWriteFile(ss, sheet)
	case "g":
		sheetWriteGraph(ss, sheet)
	case "ts":
		w := os.Stdout
		ts, err := bbk.ParseTerms(sheet)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%d terms\n", len(ts))
		for _, t := range ts {
			fmt.Fprintf(w, " - %s\n", t)
		}
	case "mk":
		tmpl := template.Must(template.New("").Parse(bbk.MakefileTemplate))

		if err := tmpl.Execute(os.Stdout, sheet); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown mode: %q", *mode)
	}
}

func sheetWriteFile(ss *bbk.SheetSet, sheet *bbk.Sheet) {
	fmt.Fprintln(os.Stdout, "\\input{../../sheet.tex}")
	fmt.Fprintln(os.Stdout, "\\sbasic")
	// just inline macros...
	for _, n := range bbk.AllNeeds(sheet, ss) {
		need := ss.Sheets[n]
		if mstring := bbk.MacrosOrEmpty(need); mstring != "" {
			fmt.Fprintf(os.Stdout, "% macros from %q\n", need.Config.Name)
			fmt.Fprintln(os.Stdout, mstring)
		}
	}
	//for _, n := range bbk.AllNeeds(sheet, ss) {
	//		fmt.Fprintf(os.Stdout, "\\input{../%s/macros.tex}\n", n)
	//	}
	//fmt.Fprintln(os.Stdout, "\\input{./macros.tex}")
	if mstring := bbk.MacrosOrEmpty(sheet); mstring != "" {
		fmt.Fprintf(os.Stdout, "% macros from %q\n", sheet.Config.Name)
		fmt.Fprintln(os.Stdout, mstring)
	}
	fmt.Fprintln(os.Stdout, "\\sstart")
	fmt.Fprintln(os.Stdout, "\\bpage\\clearpage")
	fmt.Fprintf(os.Stdout, "\\stitle{%s}{%s}\n", bbk.Title(sheet.Config.Name), sheet.Config.Name)
	//fmt.Fprintf(os.Stdout, "{\\small Needs: %s}\n", strings.Join(p.Needs, ", "))
	/*
		for _, l := range p.Lines {
			fmt.Fprintln(os.Stdout, l)
		}
	*/
	fmt.Fprintln(os.Stdout, "\\input{sheet}")
	fmt.Fprintln(os.Stdout, "\\clearpage\\gpage{graph}")
	fmt.Fprintln(os.Stdout, "\\strats")
}

func sheetWriteGraph(ss *bbk.SheetSet, sheet *bbk.Sheet) {
	var rows = [][]string{
		[]string{"name", "need"},
	}
	seen := map[*bbk.Sheet]bool{}
	q := []*bbk.Sheet{sheet}
	for len(q) > 0 {
		next := q[0]
		q = q[1:]

		if seen[next] {
			continue
		}

		seen[next] = true

		for _, n := range next.Config.Needs {
			q = append(q, ss.Sheets[n])
			rows = append(rows, []string{bbk.NodeName(next.Config.Name), bbk.NodeName(n)})
		}
	}

	// if no needs
	if len(rows) == 1 {
		// add a single entry for the node
		rows = append(rows,
			[]string{bbk.NodeName(sheet.Config.Name), ""},
		)
	}
	f, err := os.Create("graph.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	if err := w.WriteAll(rows); err != nil {
		log.Fatal(err)
	}
}

const macrosHelp = `bbk macros

This is a command to generate a macros.tex file from a sheet.lit 
inline comment.

Examples
  - bbk macros sheet.lit
`

func macrosMain(s *state) {
	if len(os.Args) < 3 {
		s.info(macrosHelp[:strings.Index(macrosHelp, "\n")])
	}
	sf, err := os.Open(os.Args[2])
	if err != nil {
		s.error("os.Open(%q): %v", os.Args[2], err)
	}
	defer sf.Close()
	sheet, err := bbk.ParseSheet(sf)
	if err != nil {
		s.error("ParseSheet %v", err)
	}
	for c := sheet.LitNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == lit.CommentNode && len(c.Data) > 5 && c.Data[:6] == "macros" {
			fmt.Print(strings.TrimSpace(c.Data[strings.Index(c.Data, "\n"):]))
		}
	}
}

const randHelp = `bbk rand

Open a random sheet.

Examples
  - bbk rand
`

func randMain(s *state) {
	if len(s.Args) < 2 {
		s.info(randHelp[:strings.Index(randHelp, "\n")])
		return
	}

	fs := flag.NewFlagSet("all", flag.ContinueOnError)
	sheetsDir := fs.String("sheets", ".", "where to find other sheets")
	if err := fs.Parse(s.Args[2:]); err != nil {
		s.error("parsing flags: %v", err)
		return
	}
	ss, err := bbk.ParseSheetSet(os.DirFS(*sheetsDir))
	if err != nil {
		s.error("error: %v", err)
	}
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(ss.Sheets) - 1
	randomInt := rand.Intn(max-min+1) + min

	var i int
	var sheet *bbk.Sheet
	for _, s := range ss.Sheets {
		if i == randomInt {
			sheet = s
			break
		}
		i += 1
	}

	cmd := exec.Command("vi", path.Join(*sheetsDir, fmt.Sprintf("%s/sheet.lit", sheet.Config.Name)))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

const allHelp = `bbk all

This is the old bbk_all command.

Examples
  - bbk all
`

func allMain(s *state) {
	fs := flag.NewFlagSet("all", flag.ContinueOnError)

	var (
		sheetsDir = fs.String("sheets", "../", "where to find other sheets")
	)

	if err := fs.Parse(s.Args[2:]); err != nil {
		s.error("parsing flags: %v", err)
		return
	}
	ss, err := bbk.ParseSheetSet(os.DirFS(*sheetsDir))
	if err != nil {
		log.Fatal(err)
	}

	rs := make(map[string]*bbk.Sheet, len(ss.Sheets))
	for _, p := range ss.Sheets {
		rs[p.Config.Name] = p
	}

	// get the template
	order := mustGetAllSheetsOrder("sheets.csv")
	orderedResults := make([]*bbk.Sheet, len(order))
	for i, n := range order {
		r, ok := rs[n]
		if !ok {
			panic(fmt.Sprintf("unknown sheet: %v", n))
		}
		orderedResults[i] = r
	}

	inputsTemplate := template.Must(
		template.New("inputs.tmpl").Funcs(
			template.FuncMap{
				"title": bbk.Title,
				"join": func(ss []string) string {
					return strings.Join(ss, ", ")
				},
			},
		).ParseFiles("inputs.tmpl"))

	if err := inputsTemplate.Execute(os.Stdout, orderedResults); err != nil {
		log.Fatal(err)
	}
}

func mustGetAllSheetsOrder(filename string) []string {
	// get the order
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("os.Open: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	var order []string
	// var entries []Entry
	var header = true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if header {
			header = false
			continue
		}

		order = append(order, record[0])
	}
	return order
}
