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

const SearchTemplateHTML = `<!DOCTYPE html>
  <head>
    <link rel="stylesheet" href="./style.css">
    <link rel="stylesheet" href="./fonts.css">
	</head>
  <body>
		<div class="page">
			<div class="content">
				<img src="../trademark.pdf" id="trademark">
				<h1>Search</h1>
				<div style="display: flex; flex-direction:column">
				<div>
					<form method="GET">
					<input style="border:none;border-bottom: 1px solid gray;font-size: 1em; padding: 3px; width: 90%; margin: 3px auto;" value="{{ .Query }}" name="query" autofocus/>
					</form>
					<div>
					{{ len .Results }} results in {{ .SearchDuration | printf "%s" }};
					<a href="./index.html">Back to index</a>
					</div>
				</div>
					{{ range $r := .Results }}
					<div tabindex="0" style="cursor:pointer; border: 1px solid gray; padding 5px; border-radius: 3px; padding: 12px; margin:3px" onclick="location.href='./sheets/{{ $r.Name}}.html'">
					<b> {{ title $r.Name }} </b> <br>
					Rank: {{ .Rank | printf "%.2f" }}; Reasons: {{ reasons . }}
					</div>
					{{end}}
				</div>
		</div>
	</div>
	</body>
	<script>
	</script>
</html>
`

var searchTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	"title":   bbk.Title,
	"reasons": reasons,
}).Parse(SearchTemplateHTML))

func reasons(sr *bbk.SearchResult) string {
	return strings.Join(sr.Reasons, ", ")
}
