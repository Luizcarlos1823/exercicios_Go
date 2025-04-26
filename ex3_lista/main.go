package main

import "fmt"

func main(){

	listNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}

	fmt.Println(CountPositivesSumNegatives(listNumbers))
	
}

func CountPositivesSumNegatives(listNumbers []int) []int {

	var count, sum int
	
	for _, number := range listNumbers {
		if number > 0 {
			count += 1
		}
		if number < 0 {
			sum += number
		}
	}
	return []int{count, sum}

  }