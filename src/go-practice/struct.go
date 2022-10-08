package main

import "fmt"

type Person struct {
	name, city string
	age        int
}

func main() {

	var p1 Person

	p1.age = 12
	p1.city = "中国"
	p1.name = "小阿悄"

	fmt.Printf("p1=%#v\n", p1)
}
