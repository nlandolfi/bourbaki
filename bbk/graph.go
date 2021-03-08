package bbk

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
)

type Entry struct {
	Name     string
	DirName  string
	Needs    []string
	NeededBy []string
}

func ParseGraph(f io.Reader) (map[string]*Entry, error) {
	r := csv.NewReader(f)
	entries := make(map[string]*Entry)
	var header = true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if header {
			header = false
			continue
		}

		name := record[0]
		entry, ok := entries[name]
		if !ok {
			entry = &Entry{
				Name:    name,
				DirName: DirName(name),
			}
			entries[name] = entry
		}
		need := record[1]
		entry.Needs = append(entry.Needs, need)
	}

	for name, e := range entries {
		for _, need := range e.Needs {
			e2, ok := entries[need]
			if !ok {
				log.Printf("name: %s, need: %s, missing entry", name, need)
				continue
			}
			e2.NeededBy = append(e2.NeededBy, name)
		}
	}

	return entries, nil
}

func NodeName(s string) string {
	pieces := strings.Split(s, "_")
	return strings.Title(strings.Join(pieces, " "))
}

func DirName(s string) string {
	pieces := strings.Split(s, " ")
	for i, p := range pieces {
		pieces[i] = strings.ToLower(p)
	}
	return strings.Join(pieces, "_")
}
