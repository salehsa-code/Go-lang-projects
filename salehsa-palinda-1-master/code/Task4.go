package main

import "fmt"
// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan <- int) {
    // TODO
    var b int 
   for i := range a {
   		b += a[i]
	   }

	res <- b 


}

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7}
    n := len(a)
    ch := make(chan int)
    go Add(a[:n/2], ch)
    go Add(a[n/2:], ch)

    r1 := <-ch 
    r2 := <-ch 

    fmt.Println(r1 , r2, "sum",  r1 + r2)
    // TODO: Get the subtotals from the channel and print their sum.
}