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

	m3 := &map[int]string{
		1: "田中",
		2: "太郎",
	}
	m4 := m3
	fmt.Printf("Address of m3: %p\n", &m3) //Address of m3: 0x14000120030
	fmt.Printf("Address of m4: %p\n", &m4) //Address of m4: 0x14000120040
	fmt.Println("m3:", m3)                 //m3: &map[1:田中 2:太郎]
	fmt.Println("m4:", m4)                 //m4: &map[1:田中 2:太郎]

	m4 = &map[int]string{
		1: "山崎",
		2: "一番",
	}
	fmt.Println("m3:", m3)                 //m3: &map[1:田中 2:太郎]
	fmt.Println("m4:", m4)                 //m4: &map[1:山崎 2:一番]
	fmt.Printf("Address of m3: %p\n", &m3) //Address of m3: 0x14000120030
	fmt.Printf("Address of m4: %p\n", &m4) //Address of m4: 0x14000120040

}
