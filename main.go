package main

import "fmt"

func main() {

	x := 1

	// xの値がコピー
	y := x

	// yの操作はxに影響しない
	y++

	fmt.Println(x)  // 1
	fmt.Println(&x) //0x140000a2008

	fmt.Println(y)  // 2
	fmt.Println(&y) //0x140000a2010

}
