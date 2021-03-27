package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"

	"bbk"
)

var (
	sheetsDir = flag.String("sheets", ".", "the sheets directory")
)

func main() {
	flag.Parse()

	results, err := bbk.ParseAll(*sheetsDir)

	if err != nil {
		log.Fatalf("bbk.ParseAll: %v", err)
	}

	log.Printf("%d sheets", len(results))

	for name, p := range results {
		tmpl := template.Must(template.New("").Parse(bbk.MakefileTemplate))

		out, err := os.Create(path.Join(*sheetsDir, name, "Makefile"))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		if err := tmpl.Execute(out, p); err != nil {
			log.Fatal(err)
		}
		out.Close()
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
	for i := 0; i < 64; i++ {
		go func(in <-chan string) {
			for name := range ch {
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

	for n := range results {
		wg.Add(1)
		ch <- n
	}

	close(ch)

	wg.Wait()
}
