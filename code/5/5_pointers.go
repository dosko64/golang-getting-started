package main

import "fmt"

func main() {
	var p *int
	i := 42
	p = &i
	fmt.Println(*p) // read i through the pointer p
	fmt.Println(p)  // value of p is a pointer
	*p = 21         // set i through the pointer p
	fmt.Println(i)
}
