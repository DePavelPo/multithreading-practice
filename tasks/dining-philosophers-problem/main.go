package main

import (
	"flag"
	"strconv"

	sl "github.com/DePavelPo/multithreading-practice/tasks/dining-philosophers-problem/solutions"
)

// go run main.go -- <philosophers number> (10 is default)
func main() {
	flag.Parse()

	gottenArg := flag.Arg(0)
	gottenArgInt, _ := strconv.Atoi(gottenArg)

	gottenArgInt = gotOrDefault(gottenArgInt, 10)

	sl.UsingWaiter(gottenArgInt)
}

func gotOrDefault(got, def int) int {
	if got == 0 {
		return def
	}
	return got
}
