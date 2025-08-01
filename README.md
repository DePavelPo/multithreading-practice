# multithreading-practice
The repository contains solutions of different tasks about multithreading in Golang

Tasks:
1. task with WaitGroup (./tasks/waitGroup-easy-task):
   - make a slice of int values (from 1 to 100)
   - find a square for each value using goroutines
   - output the result
     - Solutions:
       1. async output: finding squares and result output in goroutines using mutex to avoid data race (case number 1)
       2. sync output: finding squares in goroutines using mutex to avoid data race. Sorted output (case number 2)
       3. buffered chan: save squares to buffered channel in goroutines. Sorted output (case number 3)

2. task with producer and consumer on channel (./tasks/chan-producer-consumer):
   - producer's goroutine puts random value into channel
   - consumer's goroutine reads channel and output the square of received value
     - Solutions:
       1. using unbuffered channel and single producer and consumer (case number 1)
       2. using buffered channel, single producer and consumer, using context with cancel (case number 2)
       3. using buffered channel, few producers and single consumer, using context with cancel (case number 3)

3. task with data filter on channels (goroutines' conveyor) (./tasks/chan-data-filter):
   - first goroutine generates random number and puts it into "a" channel
   - second goroutine chooses only even numbers from "a" channel and puts them into "b" channel
   - third goroutine does x2 for each number from "b" channel and puts them into "c" channel
   - forth goroutine output numbers from "c" channel
     - Solutions:
       1. using buffered channel (case number 1)
       2. using buffered channel, handling in each goroutine if context was cancelled (case number 2)

4. task with safety counter (.tasks/safety-counter):
   - using Counter struct with value field
   - Inc() increments value, Dec() decrements value
   - GetValue() returns value
   - Calling Inc() and Dec() in different goroutines
      - Solutions:
        1. using RWMutex (case number 1)
        2. using RWMutex with Cond to prevent situations when value is negative (case number 2)
        3. using atomic instead of RWMutex and Cond (case number 3)

5. task with limit number of goroutines (semaphore and worker pool) (.tasks/goroutines-limit):
   - generate some random numbers
   - the limit number of goroutines calculate fibonacci by generated numbers concurrently
   - output calculated numbers
      - Solutions:
        1. using semaphore (case number 1)
        2. using worker pool (case number 2)

6. task with sync cache (.tasks/cache-sync):
   - there is a cache with methods get(key) (value, bool) and set(key, value)
   - needs to protect access to cache using mutex
      - Solutions:
        1. using RWMutex

7. task with future-promise realization (.tasks/future-promise):
   - the goroutine does fibonacci calculation (or other long tasks)
   - the main goroutine is blocked until the result of calculation is received
      - Solutions:
        1. using unbuffered channel (case number 1)
        2. using function that returns the channel that return the result (case number 2)
        3. using context with timeout to prevent really long proccesses (case number 3)
        4. using third-side library (case number 4)

8. task with 'Dining philosophers problem' (.tasks/dining-philosophers-problem):
   - there are N philosophers sitting at a round table
   - there is one fork between each philosopher
   - to eat, a philosopher needs two forks - left and right
   - all philosophers eats repeatedly until the dining finishes
   - the dining has a duration
      - Solutions:
        1. with one Waiter who servers philosophers