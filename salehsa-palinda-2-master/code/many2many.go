// Stefan Nilsson 2013-03-13

// This is a testbed to help you understand channels better.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgpNew := new(sync.WaitGroup)
	wgp.Add(producers)
	wgpNew.Add(consumers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch, wgpNew)
	}
	wgp.Wait() // Wait for all producers to finish.
	close(ch)
	wgpNew.Wait()
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
	wg.Done()
}

// RandomSleep waits for x ms, where x is a random number, 0 â‰¤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}


/*
What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?
-It won't work since its not possible to send to a closed channel. Result = panic: send on closed channel

What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
-Other Go routines uses uses Produce closing it before will generate an error, ohter ch won't be able to write to it

What happens if you remove the statement close(ch) completely?
-The program won't be effected

What happens if you increase the number of consumers from 2 to 4?
-The progam runs faster, it uses 4 goroutines do read the channels, it runs the costumers in paralell and reads 
the channels at the same time.  

Can you be sure that all strings are printed before the program stops?
-Like bug02 (previous task) having wgp.Wait() waits for the goroutine, ans: Yes this is because we are using an unbuffered channel.
Each producer waits for a consumer to print the value the producer wrote to the channel. This continutes until all producers are done.
.Done() is called upon when all values sent by the producers are pinted out by the conusmer  


*/
