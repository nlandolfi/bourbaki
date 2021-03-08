package main

import (
	"bbk"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	graphFile = flag.String("graph-file", "graph.csv", "the graph to check")
)

func main() {
	flag.Parse()

	f, err := os.Open(*graphFile)
	if err != nil {
		log.Fatalf("os.Open: %q, %v", *graphFile, err)
	}
	defer f.Close()

	entries, err := bbk.ParseGraph(f)
	if err != nil {
		log.Fatal(err)
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
		_, err := os.Stat(fmt.Sprintf("../sheets/%s", bbk.DirName(s)))
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
		if !seen[bbk.NodeName(f.Name())] {
			nonodes += 1
			log.Printf("%s %s", f.Name(), bbk.NodeName(f.Name()))
		}
	}

	log.Printf("%d missing nodes", nonodes)
}
