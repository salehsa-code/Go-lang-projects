// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"sync"
)

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // Make room for one unmatched send.
//	wg := new(sync.WaitGroup)
	var wg sync.WaitGroup			
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %sâ€™s message.\n", name)
	default:
		// There was no pending send operation.
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// Wait for someone to receive my message.
	}
	wg.Done()
}
/*
What happens if you remove the go-command from the Seek call in the main function?
Since it's buffered nothing will happen 

What happens if you switch the declaration wg := new(sync.WaitGroup)
 to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?

 A deadlock occurs, 

 What happens if you remove the buffer on the channel match?
 Deadlock, message being sent without being read
 
 What happens if you remove the default-case from the case-statement in the main function?
 there will always be a sending operation due to odd number of people 

*/