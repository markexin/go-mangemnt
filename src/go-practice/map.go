package main

import "fmt"

func main() {
	sourceMap := map[string]int{
		"username": 13,
		"password": 8,
	}
	fmt.Println(sourceMap)
	v, ok := sourceMap["username"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
