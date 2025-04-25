package main

import "fmt"

func main(){
	numbers := []int{1, -4, 7, 12, -2}
	var sum int

	for _, number := range numbers{
		if number >= 0 {
			sum += number
		}		
	}
	fmt.Println(sum)
}
