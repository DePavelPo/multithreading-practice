package main

import (
	"flag"
	"log"
	"strconv"

	sl "github.com/DePavelPo/multithreading-practice/tasks/waitGroup-easy-task/solutions"
)

// go run main.go -- <case number>
func main() {
	flag.Parse()

	gottenFlag := flag.Arg(0)

	solutionNum, err := strconv.Atoi(gottenFlag)
	if err != nil {
		log.Fatal("unexpected flag format, must be int")
	}

	switch solutionNum {
	case 1:
		sl.RandomOutput()
	case 2:
		sl.SyncOutput()
	case 3:
		sl.BufferedChan()
	}
}
