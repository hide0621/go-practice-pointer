package main

import "fmt"

func main() {
	i := 1
	fmt.Printf("main関数のiのメモリアドレスは%v\n", &i) //0x14000126008
	// 引数の値渡し
	comfirmAddress(i) //0x14000126010
}

func comfirmAddress(i int) {
	fmt.Printf("サブ関数のiのメモリアドレスは%v\n", &i)
}
