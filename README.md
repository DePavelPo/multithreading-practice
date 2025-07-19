# multithreading-practice
The repository contains solutions of different tasks about multithreading in Golang

Tasks:
1. task with WaitGroup (./tasks/waitGroup-easy-task):
   - make a slice of int values (from 1 to 100)
   - find a square for each value using goroutines
   - output the result
     - Solutions:
       1. async output: finding squares and result output in goroutines using mutex to avoid data race
       2. sync output: finding squares in goroutines using mutex to avoid data race. Sorted output
       3. buffered chan: save squares to buffered channel in goroutines. Sorted output

2. task with producer and consumer on channel (./tasks/chan-producer-consumer):
   - producer's goroutine puts random value into channel
   - consumer's goroutine reads channel and output the square of received value
     - Solutions:
       1. using unbuffered channel and single producer and consumer
