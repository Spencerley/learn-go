package arrays

import "fmt"

var array_1 = [3]int{1, 2, 3} // length is defined
var array_2 = [...]int{4, 5, 6} // length is inferred
var array_3 = [5]string{0:"a", 1:"b", 2:"c"}

func Arrays() {
	fmt.Println(array_1)
	fmt.Println(array_2)
}

func PriceChange() {
	prices := [3]int{10,20,30}
  
	prices[1] = 50
	fmt.Println(prices, array_3)
  }
