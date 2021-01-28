package main 

import (
		"fmt"
		"time"
		"testing")




func Remind(text string, delay time.Duration) {
		
		for{
			time.Sleep(delay)
			fmt.Println("Klockan är", time.Now().Format("3:04:05PM"), text)

		}


	
}



func main(){
	
	go Remind("Dags att äta", time.Hour * 3 )
	go Remind("Dags att arbeta", time.Hour * 5 )
	go Remind("Dags att sova", time.Hour * 24 )
	
	select{}

}
