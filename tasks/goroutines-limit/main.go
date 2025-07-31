package main

import (
	"flag"
	"log"
	"strconv"

	sl "github.com/DePavelPo/multithreading-practice/tasks/goroutines-limit/solutions"
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
		sl.SemaphoreUsing(100, 5)
	case 2:
		sl.WorkerPoolUsing(100, 5)
	}
}
