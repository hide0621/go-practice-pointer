package main

import "fmt"

func main() {

	var hoge string = "テストです"

	var pointer *string = &hoge

	// "テストです"
	fmt.Println(hoge)

	// アドレス "0x000096210"
	fmt.Println(pointer)
}
