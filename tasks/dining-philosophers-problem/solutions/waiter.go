package solutions

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Request struct {
	PhilosopherID int
	ResponseChan  chan bool
}

// waiter works with philosophers requests
// decides who can use forks
func waiter(ctx context.Context, forks int, wg *sync.WaitGroup, request chan Request, clear <-chan int) {
	var mu sync.Mutex
	forksAvailable := make(map[int]bool, forks)
	for i := 0; i < forks; i++ {
		forksAvailable[i] = true
	}
	defer wg.Done()

	var queue []Request
	wgLocal := sync.WaitGroup{}

	wgLocal.Add(1)
	// work with philosophers' requests
	go func() {
		defer wgLocal.Done()
		for {
			select {
			case task := <-request:
				// choose left and right forks from philosopher
				leftfork := task.PhilosopherID
				// if there are 10 philosophers: (9 + 1) % 10 = 0 - fork with id = 0
				rightfork := (task.PhilosopherID + 1) % forks

				// if both forks are free - waiter can give a permission
				mu.Lock()
				if forksAvailable[leftfork] && forksAvailable[rightfork] {
					forksAvailable[leftfork] = false
					forksAvailable[rightfork] = false

					task.ResponseChan <- true
				} else {
					// waiter will check if forks become available
					queue = append(queue, task)
				}
				mu.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()

	wgLocal.Add(1)
	// work with alerts that philosopher finished the meal
	go func() {
		defer wgLocal.Done()
		for {
			select {
			case philosopherID := <-clear:
				// choose left and right forks from philosopher
				leftfork := philosopherID
				// if there are 10 philosophers: (9 + 1) % 10 = 0 - fork with id = 0
				rightfork := (philosopherID + 1) % forks

				// waiter can free forks if the philosopher finishes the meal
				mu.Lock()
				forksAvailable[leftfork] = true
				forksAvailable[rightfork] = true
				mu.Unlock()

				mu.Lock()
				// check if waiter can serve some waited philosophers
				newQueue := queue[:0]
				for _, request := range queue {
					leftfork := request.PhilosopherID
					rightfork := (request.PhilosopherID + 1) % forks
					if forksAvailable[leftfork] && forksAvailable[rightfork] {
						forksAvailable[leftfork] = false
						forksAvailable[rightfork] = false

						request.ResponseChan <- true
					} else {
						newQueue = append(newQueue, request)
					}
				}
				queue = newQueue
				mu.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()

	wgLocal.Wait()
}

func UsingWaiter(philosophers int, diningDuration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), diningDuration)
	defer cancel()

	forksRequest := make(chan Request, philosophers)
	mealFinish := make(chan int, philosophers)

	// Work with philosophers
	wgPhilosophers := &sync.WaitGroup{}
	for i := 0; i < philosophers; i++ {
		wgPhilosophers.Add(1)

		go func(idx int) {
			defer wgPhilosophers.Done()

			responseChan := make(chan bool)
			defer close(responseChan)

			for {
				fmt.Printf("Philosopher %d start\n", i)
				forksRequest <- Request{PhilosopherID: idx, ResponseChan: responseChan}

				select {
				// philosopher waits until waiter gives a permission to eat
				case canEat := <-responseChan:
					if canEat {
						time.Sleep(time.Second)
						mealFinish <- idx
					}

					fmt.Printf("Philosopher %d finish\n", idx)
					time.Sleep(time.Second)
				case <-ctx.Done():
					return
				}
			}

		}(i)
	}

	// Wait until philosophers finish their meal
	// And then close channels that waiter works with
	go func() {
		wgPhilosophers.Wait()
		close(forksRequest)
		close(mealFinish)
	}()

	wgWaiter := &sync.WaitGroup{}
	wgWaiter.Add(1)
	// Work with waiter
	go waiter(ctx, philosophers, wgWaiter, forksRequest, mealFinish)
	wgWaiter.Wait()
}
