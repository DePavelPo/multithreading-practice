package main

import (
	"flag"
	"log"
	"strconv"

	sl "github.com/DePavelPo/multithreading-practice/tasks/future-promise/solutions"
)

// go run main.go -- <case number>
func main() {
	flag.Parse()

	gottenArg := flag.Arg(0)
	gottenArgInt, err := strconv.Atoi(gottenArg)
	if err != nil {
		log.Fatal("unexpected flag format, must be int")
	}

	switch gottenArgInt {
	case 1:
		sl.FuturePromise()
	case 2:
		sl.FutureWithReturnedChan()
	case 3:
		sl.FutureWithTimeout()
	case 4:
		sl.FutureWithLibrary()
	}
}
