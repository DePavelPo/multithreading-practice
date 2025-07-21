package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	sl "github.com/DePavelPo/multithreading-practice/tasks/chan-data-filter/solutions"
)

// go run main.go -- <case number>
// with time measure: go run main.go -- <case number> -t
func main() {
	flag.Parse()

	gottenArg := flag.Arg(0)
	gottenArgInt, err := strconv.Atoi(gottenArg)
	if err != nil {
		log.Fatal("unexpected flag format, must be int")
	}

	withTimeMeasure := false
	gottenSecArg := flag.Arg(1)
	if gottenSecArg == "-t" {
		withTimeMeasure = true
	}

	switch gottenArgInt {
	case 1:
		if withTimeMeasure {
			measureTime(func() {
				sl.DataFilter(200)
			})
		} else {
			sl.DataFilter(200)
		}
	case 2:
		if withTimeMeasure {
			measureTime(func() {
				sl.ContextCancelHandling(200)
			})
		} else {
			sl.ContextCancelHandling(200)
		}
	}
}

func measureTime(f func()) {
	start := time.Now()
	f()
	fmt.Printf("func finished for %v", time.Since(start))
}
