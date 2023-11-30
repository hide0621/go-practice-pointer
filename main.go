package main

import "fmt"

func main() {
	m1 := map[int]string{
		1: "山田",
		2: "鈴木",
	}

	fmt.Printf("Address of m1: %p\n", &m1) //0x14000114018
	fmt.Println(m1[1])                     //山田

	m2 := m1

	fmt.Printf("Address of m2: %p\n", &m2) //0x14000114028
	fmt.Println("key1 of m2:", m2[1])      //山田
	fmt.Println("key1 of m1:", m1[1])      //山田

	m2[1] = "John" //m1[1]の値も書き変わる

	fmt.Println("key1 of m2:", m2[1])      //John
	fmt.Println("key1 of m1:", m1[1])      //John
	fmt.Printf("Address of m1: %p\n", &m1) //0x14000114018
	fmt.Printf("Address of m2: %p\n", &m2) //0x14000114028

	m2 = make(map[int]string)
	m2[1] = "x"                            //m1[1]の値は変わらない
	fmt.Printf("Address of m2: %p\n", &m2) //0x14000114028
	fmt.Println("key1 of m1:", m1[1])      //John
	fmt.Println("key1 of m2:", m2[1])      //x

	m2 = map[int]string{
		1: "y",
		2: "z",
	}
	fmt.Printf("Address of m2: %p\n", &m2) //0x14000114028
	fmt.Println("key1 of m1:", m1[1])      //John //m1[1]の値は変わらない
	fmt.Println("key1 of m2:", m2[1])      //y

	m2[1] = "John" //m1[1]の値も書き変わる

	fmt.Println("key1 of m2:", m2[1])      //John
	fmt.Println("key1 of m1:", m1[1])      //John
	fmt.Printf("Address of m1: %p\n", &m1) //0x14000114018
	fmt.Printf("Address of m2: %p\n", &m2) //0x14000114028

}
