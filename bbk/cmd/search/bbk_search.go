package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"bbk"
)

var (
	sheetsDir = flag.String("sheets-dir", "../../sheets", "the sheets directory")
)

func main() {
	flag.Parse()

	results, err := bbk.ParseAll(*sheetsDir)
	if err != nil {
		log.Fatal(err)
	}

	s := bbk.NewSearcher(results)

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
			fmt.Printf("  %d @ (%.2f) %s: %s\n", i+1, r.Rank, r.Name, strings.Join(r.Reasons, ", "))
		}
		fmt.Printf("> ")
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}
}

type Config struct {
	MaxResults int
}
