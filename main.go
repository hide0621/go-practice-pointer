package main

import "fmt"

func main() {

	x := 1

	var y *int = &x

	fmt.Println(x)  //1
	fmt.Println(*y) //1

	*y++

	fmt.Println(x)  // 2
	fmt.Println(&x) //0x140000180a0

	fmt.Println(*y) // 2
	fmt.Println(&y) //0x1400000e028

}
