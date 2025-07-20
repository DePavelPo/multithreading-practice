package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	sl "github.com/DePavelPo/multithreading-practice/tasks/chan-producer-consumer/solutions"
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
	gottenSecondArg := flag.Arg(1)
	if gottenSecondArg == "-t" {
		withTimeMeasure = true
	}

	switch gottenArgInt {
	case 1:
		if withTimeMeasure {
			measureTime(sl.UnbufferedChannel)
		} else {
			sl.UnbufferedChannel()
		}
	case 2:
		if withTimeMeasure {
			measureTime(sl.BufferedChannel)
		} else {
			sl.BufferedChannel()
		}
	case 3:
		if withTimeMeasure {
			measureTime(sl.FewProducers)
		} else {
			sl.FewProducers()
		}
	}
}

func measureTime(f func()) {
	start := time.Now()
	f()
	fmt.Printf("func finished for %v", time.Since(start))
}
