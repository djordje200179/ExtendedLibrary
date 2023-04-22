package sources

import (
	"bufio"
	"github.com/djordje200179/extendedlibrary/mapreduce"
	"github.com/djordje200179/extendedlibrary/misc"
	"log"
	"os"
)

func NewFileSource(path string) mapreduce.Source[int, string] {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	source := make(chan misc.Pair[int, string], 1)

	go func() {
		lineIndex := 0
		for scanner.Scan() {
			source <- misc.Pair[int, string]{lineIndex, scanner.Text()}
			lineIndex++
		}
		close(source)
	}()

	return source
}
