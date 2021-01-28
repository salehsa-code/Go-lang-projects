
// Loops and Functions

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x)
	
	
/*	for j:=0; j<11; j++{	
	z -= (z*z - x) / (2*z)
	fmt.Println(z)

	}
*/
	
	z2 := float64(1)
	delta := float64(1)
	for delta >= 1e-10{
	
		z -= (z*z - x) / (2*z)
		delta = math.Abs(z2 - z) 
		z2 = z	
		

		
		fmt.Println(z)
		
	}


	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("math.sqrt ", math.Sqrt(2))
}


//Slices 



package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    image := make([][]uint8, dy)
    
	for x := range(image) {
        image[x] = make([]uint8, dx)
    }
    for x := range(image) {
        for y := range(image[x]) {
            image[x][y] = uint8 (x^y)
        }
    }
    return image
}

func main() {
    pic.Show(Pic)
}

//Maps




package main

import (
	"golang.org/x/tour/wc"
	"strings"
//	"fmt"

)

func WordCount(s string) map[string]int {
	x := strings.Fields(s) 
//	fmt.Println(x)
	var words map[string]int = make(map[string]int)
	
	for i := range (x){
		count := 0
		for j:=range (x){
			if x[i] == x[j]{
			count++
		}
		}
		words[x[i]] = count
		
	}

	
	return words
}

func main() {
	wc.Test(WordCount)
}



//Fibbonacci Closure

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := int(-1)
	y := int(1)
	z := 0
	return func() int{
		z = x + y
		x = y
		y = z 
		
	
		return z  	
	}
	
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}











