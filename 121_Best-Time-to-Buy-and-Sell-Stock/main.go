package main

import "fmt"

// 121. Best Time to Buy and Sell Stock
func main() {
	prices := []int{7, 8, 5, 3, 6, 4}

	maxProfit := 0
	minPrice := prices[0]

	for _, i := range prices {
		if minPrice > i {
			minPrice = i
			fmt.Println(minPrice)
		} else if maxProfit < i-minPrice {
			maxProfit = i - minPrice
		}
	}
	fmt.Println(maxProfit)
}
