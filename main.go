package main

import "fmt"

func main() {

	var hoge string = "テストです"
	var pointer *string = &hoge

	// "テストです"
	fmt.Println(*pointer)

}
