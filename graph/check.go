package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Entry struct {
	Name  string
	Needs []string
}

func main() {
	f, err := os.Open("graph.csv")
	if err != nil {
		log.Fatal("os.Open: %v", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	entries := make(map[string]*Entry)
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

		name := record[0]
		entry, ok := entries[name]
		if !ok {
			entry = &Entry{
				Name: name,
			}
			entries[name] = entry
		}
		need := record[1]
		entry.Needs = append(entry.Needs, need)
	}

	seen := make(map[string]bool)
	for _, e := range entries {
		seen[e.Name] = true
		for _, n := range e.Needs {
			seen[n] = true
		}
	}

	var nodirectories int
	for s := range seen {
		_, err := os.Stat(fmt.Sprintf("../sheets/%s", dirname(s)))
		if os.IsNotExist(err) {
			nodirectories += 1
			log.Print(s)
		} else if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("%d missing sheet directories", nodirectories)

	files, err := ioutil.ReadDir("../sheets/")
	if err != nil {
		log.Fatal(err)
	}

	var nonodes = 0

	for _, f := range files {
		if !seen[nodename(f.Name())] {
			nonodes += 1
			log.Printf("%s %s", f.Name(), nodename(f.Name()))
		}
	}

	log.Printf("%d missing nodes", nonodes)

}

func nodename(s string) string {
	pieces := strings.Split(s, "_")
	return strings.Title(strings.Join(pieces, " "))
}

func dirname(s string) string {
	pieces := strings.Split(s, " ")
	for i, p := range pieces {
		pieces[i] = strings.ToLower(p)
	}
	return strings.Join(pieces, "_")
}
