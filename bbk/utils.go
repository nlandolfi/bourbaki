package bbk

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// from https://opensource.com/article/18/6/copying-files-go
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

var shortWords = map[string]bool{
	"of":  true,
	"the": true,
	"a":   true,
	"an":  true,
}

func Title(s string) string {
	pieces := strings.Split(s, "_")
	pieces[0] = strings.Title(pieces[0])
	for i := 1; i < len(pieces); i++ {
		if shortWords[pieces[i]] {
			continue
		}
		pieces[i] = strings.Title(pieces[i])
	}
	return strings.Join(pieces, " ")
}
