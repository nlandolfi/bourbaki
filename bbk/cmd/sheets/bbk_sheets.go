package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"sync"

	"bbk"
)

var (
	sheetsDir = flag.String("sheets", ".", "the sheets directory")
	startAt   = flag.String("start-at", "", "sheet to start compiling at (alphabetical)")
)

func main() {
	flag.Parse()

	rs, err := bbk.ParseAll(*sheetsDir)

	if err != nil {
		log.Fatalf("bbk.ParseAll: %v", err)
	}
	results := make(map[string]*bbk.ParseResult, len(rs))
	for _, p := range rs {
		results[p.Name] = p
	}

	if len(results) == 0 {
		log.Printf("warning: no sheets parsed")
	}

	names := make([]string, 0, len(results))
	for n := range results {
		names = append(names, n)
	}
	sort.Strings(names)

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
