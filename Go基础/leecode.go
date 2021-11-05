package main

import "fmt"

func plusOne(digits [3]int) {
	currrent := &digits
	cache := ""
	for v := range currrent {
		cache = cache + string(v)
		fmt.Printf(string(v))
	}
}

func main() {
	nums := [...]int{4, 5, 5}
	plusOne(nums)
}
