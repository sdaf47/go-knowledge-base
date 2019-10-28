package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"sync"
	"time"
)

var filesCntF *int
var elementsCnt *int

func init() {
	filesCntF = flag.Int("f", 5, "number of files for creating")
	elementsCnt = flag.Int("e", 1000000, "number elements")
}

func main() {
	flag.Parse()

	logrus.
		WithField("files", *filesCntF).
		WithField("elements", *elementsCnt).
		WithField("top", "For usage try -help").
		Info("Start")

	set := make([]int, *elementsCnt)
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= *elementsCnt; i++ {
		set[i-1] = int(rand.Intn(*elementsCnt))
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	logrus.WithField("path", dir).Info("base data dir")

	filesCnt := *filesCntF
	wg := sync.WaitGroup{}
	for i := 0; i < filesCnt; i++ {
		wg.Add(1)
		start := len(set) / filesCnt * i
		end := start + len(set)/filesCnt

		fileName := fmt.Sprintf("data_%d.txt", i)
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		go func(file *os.File, start, end int) {
			for i := start; i < end; i += 2 {
				_, err := fmt.Fprintf(file, "%d %d\n", set[i], set[i+1])
				if err != nil {
					panic(err)
				}
			}
			logrus.WithField("name", fileName).Info("new file created")
			wg.Done()
		}(file, start, end)
	}
	wg.Wait()
	logrus.Info("all files created")
}
