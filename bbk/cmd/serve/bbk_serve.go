package main

import (
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"bbk"
	"spin/infra/audit"
	"spin/infra/key"
)

var (
	staticDir  = flag.String("static-dir", "/static", "directory of static files")
	localAudit = flag.Bool("local-audit", false, "only locally log audit")
	auditURL   = flag.String("audit-url", "https://audit-5rcmpbrboq-uw.a.run.app", "audit service to use")
)

var noKey = &key.Key{
	Name: "bbk-no-key",
	Type: "bbk-no-key",
}

func main() {
	flag.Parse()

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

	f, err := os.Open(path.Join(*staticDir, "results.gob"))
	if err != nil {
		log.Fatal(err)
	}
	var sd bbk.SearchData
	if err := gob.NewDecoder(f).Decode(&sd); err != nil {
		log.Fatal(err)
	}
	var results = make(map[string]*bbk.ParseResult, len(sd.Results))
	for _, r := range sd.Results {
		results[r.Name] = r
	}
	s := &searcher{sd.Version, bbk.NewSearcher(results), as}

	m := http.NewServeMux()
	m.HandleFunc("/search", s.search)
	m.HandleFunc("/search_json", s.search_json)

	fs := http.FileServer(http.Dir(*staticDir))
	m.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)

		// LOGGING
		data := map[string]string{
			"remote-addr": r.RemoteAddr,
			"url":         r.URL.String(),
		}
		for k, vs := range r.Header {
			data[k] = strings.Join(vs, ",,")
		}
		s.AuditLog.Put("get", noKey, data)
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
	}
	for k, vs := range r.Header {
		data[k] = strings.Join(vs, ",,")
	}
	s.AuditLog.Put("search", noKey, data)
}

func (s *searcher) search(w http.ResponseWriter, r *http.Request) {
	sd := s.data(r)
	searchTemplate := template.Must(
		template.New("search.tmpl").Funcs(
			template.FuncMap{
				"title":   bbk.Title,
				"reasons": reasons,
			},
		).ParseFiles(path.Join(*staticDir, "search.tmpl")))
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
