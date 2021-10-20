package main

import "fmt"

func main() {
	arr := [...]int{6, 2, 4, 9, 8, 3}
	sum := 0
	for _, a := range arr {
		sum += a
	}
	fmt.Println(sum)
}
