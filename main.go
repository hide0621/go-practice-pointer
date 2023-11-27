package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func NewPersonValue(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	//ポインタ型のインスタンスを作成
	personPointer := NewPersonPointer("山田", "太郎", 30)

	fmt.Println("First Name (Pointer):", personPointer.FirstName)
	fmt.Println("Last Name (Pointer):", personPointer.LastName)
	fmt.Println("Age (Pointer):", personPointer.Age)

	personPointer1 := NewPersonPointer("近藤", "勇", 21)
	personPointer = personPointer1
	fmt.Println("First Name (Pointer):", personPointer.FirstName) //近藤

	p := &Person{
		FirstName: "斉藤",
		LastName:  "一",
		Age:       35,
	}
	fmt.Println(p)           //&{斉藤 一 35}
	fmt.Println(p.FirstName) //斉藤

	p.FirstName = "一条"
	fmt.Println(p.FirstName) //一条

	//値型のインスタンスを作成
	personValue := NewPersonValue("佐藤", "二郎", 25)

	fmt.Println("\nFirst Name (Value):", personValue.FirstName)
	fmt.Println("Last Name (Value):", personValue.LastName)
	fmt.Println("Age (Value):", personValue.Age)

	personValue1 := NewPersonValue("土方", "歳三", 17)
	personValue = personValue1
	fmt.Println("First Name (Value):", personValue.FirstName) //土方

	v := Person{
		FirstName: "沖田",
		LastName:  "総司",
		Age:       18,
	}
	fmt.Println(v)           //{沖田 総司 18}
	fmt.Println(v.FirstName) //沖田

	v.FirstName = "轟"
	fmt.Println(v.FirstName) //轟

}
