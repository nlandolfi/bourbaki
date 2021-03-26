package main

import (
	"flag"
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

		if err := tmpl.Execute(os.Stdout, p); err != nil {
			log.Fatal(err)
		}
		out, err := os.Create(path.Join(*sheetsDir, name, "Makefile"))
		if err != nil {
			log.Fatalf("os.Create: %v", err)
		}
		if err := tmpl.Execute(out, p); err != nil {
			log.Fatal(err)
		}
		out.Close()
	}

	var wg sync.WaitGroup
	errs := make(chan error, len(results))
	ch := make(chan string, len(results))

	// these are the workers
	for i := 0; i < 10; i++ {
		go func(in <-chan string) {
			for name := range ch {
				c := exec.Command("make")
				c.Dir = name
				_, err := c.Output()
				if err != nil {
					errs <- err
				}
				wg.Done()
			}
		}(ch)
	}

	for n := range results {
		ch <- n
	}

	close(ch)

	wg.Wait()

	select {
	case err := <-errs:
		log.Fatalf("first error: %v", err)
	default:
	}

}
