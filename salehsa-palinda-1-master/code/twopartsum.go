package main

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {
	// TODO: return the sum of the numbers on the given channel
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	// TODO: Get the subtotals from the channel and print their sum.
}
