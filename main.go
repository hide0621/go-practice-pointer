package main

import "fmt"

func main() {

	a := 42
	b := &a

	fmt.Println(a) //42
	fmt.Println(b) //0x140000a4008

	fmt.Println(*b) //42
	fmt.Println(&b) //0x1400009e018
}
