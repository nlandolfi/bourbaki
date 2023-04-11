package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"bbk"

	"github.com/nlandolfi/lit"
)

var (
	sheetsDir = flag.String("sheets-dir", "../../sheets", "the sheets directory")
	staticDir = flag.String("static-dir", "./static", "the static directory")
)

func main() {
	flag.Parse()

	ss, err := bbk.ParseSheetSet(os.DirFS(*sheetsDir))
	if err != nil {
		log.Fatal(err)
	}

	var results = make(map[string]*bbk.SheetResult, len(ss.Sheets))
	var b bytes.Buffer
	for _, r := range ss.Sheets {
		b.Reset()
		lit.WriteTex(&b, r.LitNode, &lit.WriteOpts{Prefix: "", Indent: " "})
		ts, err := bbk.ParseTerms(r)
		if err != nil {
			log.Fatal(err)
		}
		results[r.Config.Name] = &bbk.SheetResult{
			Config: r.Config,
			Terms:  ts,
			RawTex: b.String(),
		}
	}

	s := bbk.NewSearcher(results, *staticDir)

	var c Config
	c.MaxResults = 3

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	for scanner.Scan() {
		rq := scanner.Text()
		if strings.HasPrefix(rq, "%!") {
			if strings.HasPrefix(rq, "%!max:") {
				i, err := strconv.ParseInt(strings.TrimPrefix(rq, "%!max:"), 10, 64)
				if err != nil {
					log.Print(err)
					continue
				}
				c.MaxResults = int(i)
			}
			fmt.Printf("config: %+v\n", c)
			continue
		}
		rrq := bbk.RewriteQuery(rq)
		rs := s.Search(rrq)
		fmt.Printf("  %% %q -> %q; showing %d/%d results\n", rq, rrq, int(math.Min(float64(len(rs)), float64(c.MaxResults))), len(rs))
		for i := 0; i < len(rs) && i < c.MaxResults; i++ {
			r := rs[i]
			fmt.Printf("  %d @ (%.2f) %s: %s\n", i+1, r.Rank, r.Config.Name, strings.Join(r.Reasons, ", "))
		}
		fmt.Printf("> ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	MaxResults int
}
