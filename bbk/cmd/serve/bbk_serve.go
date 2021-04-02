package main

import (
	"encoding/gob"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"bbk"
)

var (
	staticDir = flag.String("static-dir", "/static", "directory of static files")
)

func main() {
	flag.Parse()
	f, err := os.Open(path.Join(*staticDir, "results.gob"))
	if err != nil {
		log.Fatal(err)
	}
	var results map[string]*bbk.ParseResult
	if err := gob.NewDecoder(f).Decode(&results); err != nil {
		log.Fatal(err)
	}
	s := &searcher{bbk.NewSearcher(results)}

	m := http.NewServeMux()
	m.HandleFunc("/search", s.search)
	m.Handle("/", http.FileServer(http.Dir(*staticDir)))
	log.Fatal(http.ListenAndServe(":8080", m))
}

type searcher struct {
	*bbk.Searcher
}

func (s *searcher) search(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	searchTemplate := template.Must(
		template.New("search.tmpl").Funcs(
			template.FuncMap{
				"title":   bbk.Title,
				"reasons": reasons,
			},
		).ParseFiles(path.Join(*staticDir, "search.tmpl")))

	var sd SearchData
	sd.Query = r.FormValue("query")
	sd.RewrittenQuery = bbk.RewriteQuery(sd.Query)
	if sd.RewrittenQuery != "" {
		sd.Results = s.Search(sd.RewrittenQuery)
		sd.Searched = true
	}
	sd.SearchDuration = time.Now().Sub(start)

	if err := searchTemplate.Execute(w, &sd); err != nil {
		http.Error(w, "internal", 500)
		log.Fatal(err)
	}

}

type SearchData struct {
	Query          string
	RewrittenQuery string
	Results        []*bbk.SearchResult
	Searched       bool
	SearchDuration time.Duration
}

func reasons(sr *bbk.SearchResult) string {
	return strings.Join(sr.Reasons, ", ")
}
