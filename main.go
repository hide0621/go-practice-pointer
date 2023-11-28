package main

import "fmt"

func main() {

	a := 42
	b := &a

	fmt.Println(a) //42
	fmt.Println(b) //0x14000126008

	fmt.Println(*b) //42
	fmt.Println(&b) //42
}
