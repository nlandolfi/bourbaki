package bbk

import (
	"log"
	"sort"
	"strings"
)

type SheetResult struct {
	Config SheetConfig
	Terms  []string
	RawTex string
}

type SearchData struct {
	Results []*SheetResult
	Version string
}

type SearchResult struct {
	*SheetResult
	Reasons []string
	Ranks   []float64
	Rank    float64
}

func NewSearcher(results map[string]*SheetResult, staticDir string) *Searcher {
	names := make([]string, 0, len(results))
	for n := range results {
		names = append(names, n)
	}
	sort.Strings(names)

	var s = Searcher{
		names:     names,
		sheets:    results,
		StaticDir: staticDir,
	}
	s.spacenames = make([]string, len(names))
	for i, n := range s.names {
		s.spacenames[i] = strings.ReplaceAll(n, "_", " ")
	}
	s.terms = make(map[string][]*SheetResult)
	for _, r := range results {
		for _, t := range r.Terms {
			s.terms[t] = append(s.terms[t], r)
		}
	}
	return &s
}

type Searcher struct {
	StaticDir  string
	names      []string
	spacenames []string
	sheets     map[string]*SheetResult
	terms      map[string][]*SheetResult
}

func RewriteQuery(q string) string {
	return strings.ToLower(q)
}

func addresult(rs map[*SheetResult]*SearchResult, pr *SheetResult, reason string, rank float64) {
	if pr == nil {
		log.Fatal("attempting to add nil parse result")
	}
	sr, ok := rs[pr]
	if !ok {
		sr = &SearchResult{
			SheetResult: pr,
		}
		rs[pr] = sr
	}

	sr.Reasons = append(sr.Reasons, reason)
	sr.Ranks = append(sr.Ranks, rank)
	sr.Rank += rank
	return
}

func (s *Searcher) Search(q string) []*SearchResult {
	rs := make(map[*SheetResult]*SearchResult)
	for i, name := range s.spacenames {
		if q == name {
			addresult(rs, s.sheets[s.names[i]], "name exact match", 1)
			continue
		}
		if strings.Contains(name, q) {
			addresult(rs, s.sheets[s.names[i]], "name contains query",
				float64(len(q))/float64(len(name)),
			)
		}
	}
	for term, prs := range s.terms {
		if q == term {
			for _, pr := range prs {
				addresult(rs, pr, "term exact match", 1)
			}
			continue
		}
		if strings.Contains(term, q) {
			for _, pr := range prs {
				addresult(rs, pr, "term contains query",
					float64(len(q))/float64(len(term)),
				)
			}
		}
	}
	for _, pr := range s.sheets {
		if strings.Contains(strings.ToLower(pr.RawTex), q) {
			addresult(rs, pr, "body contains query",
				float64(len(q))/float64(len(pr.RawTex)),
			)
		}
	}
	results := make([]*SearchResult, 0, len(rs))
	for _, sr := range rs {
		results = append(results, sr)
	}
	sort.Sort(ByRank(results))
	return results
}

type ByRank []*SearchResult

func (br ByRank) Len() int           { return len(br) }
func (br ByRank) Less(i, j int) bool { return br[i].Rank > br[j].Rank }
func (br ByRank) Swap(i, j int)      { br[i], br[j] = br[j], br[i] }
