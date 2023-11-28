package main

import "fmt"

func main() {

	var p *int
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
	// p is nil

	switch {
	case p == nil:
		fmt.Println("p is nil")
	default:
		fmt.Println("p is not nil")
	}
	// p is nil

}
