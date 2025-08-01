package main

import (
	"flag"
	"strconv"
	"time"

	sl "github.com/DePavelPo/multithreading-practice/tasks/dining-philosophers-problem/solutions"
)

// go run main.go -- <philosophers number> <dining duration>
// philosophers number = 10 as default
// dining duration is in seconds, 10 seconds as a default
func main() {
	flag.Parse()

	gottenArg1 := flag.Arg(0)
	gottenArgInt1, _ := strconv.Atoi(gottenArg1)
	gottenArgInt1 = gotOrDefault(gottenArgInt1, 10)

	gottenArg2 := flag.Arg(1)
	gottenArgInt2, _ := strconv.Atoi(gottenArg2)
	gottenArgInt2 = gotOrDefault(gottenArgInt2, 10)

	sl.UsingWaiter(gottenArgInt1, time.Second*time.Duration(gottenArgInt2))
}

func gotOrDefault(got, def int) int {
	if got == 0 {
		return def
	}
	return got
}
