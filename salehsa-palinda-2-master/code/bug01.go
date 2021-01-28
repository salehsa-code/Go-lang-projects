package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	go func(){ ch <- "Hello world!"}()
	fmt.Println(<-ch)
}
/* The deadlock was fixed by adding go func() 
   The program gets stuck on the channel send operation 
   waiting forever for someone to read the value
   that's why a new go routine needs to be started 	

*/ 