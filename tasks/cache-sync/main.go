package main

import (
	sl "github.com/DePavelPo/multithreading-practice/tasks/cache-sync/solutions"
)

// go run main.go
func main() {
	sl.SyncWithMutex(50, 1000)
}
