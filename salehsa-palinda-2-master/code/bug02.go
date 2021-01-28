package main

import ("fmt"
		 "sync"
		 "time")

var wg sync.WaitGroup
// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {
	ch := make(chan int)
	
	go Print(ch)
	wg.Add(1) // set the number of goroutines to wait for
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	
	close(ch)
	wg.Wait() // waits for the goroutine 
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond)
		fmt.Println(n)
		
	}
	wg.Done()
}


/* source http://yourbasic.org/golang/wait-for-goroutines-waitgroup/
   The channell closes down berfore Pint() has printet out all numbers
   To fix this wg.Add was used to wait for the go routines

*/
