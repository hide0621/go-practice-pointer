package main

import "fmt"

type Food struct {
	ID          int
	Name        string
	Description string
}

func main() {
	fmt.Println(MakeResult()["うまいもの"])  //&{1 焼肉 すこぶるうまい}
	fmt.Println(*MakeResult()["うまいもの"]) //{1 焼肉 すこぶるうまい}
}

// エラー：cannot assign to struct field result["うまいもの"].Description in map
// func MakeResult() map[string]Food {
// 	result := make(map[string]Food)
// 	result["うまいもの"] = Food{
// 		ID:   1,
// 		Name: "焼肉",
// 	}
// 	result["うまいもの"].Description = "すこぶるうまい"

// 	return result
// }

func MakeResult() map[string]*Food {
	result := make(map[string]*Food)
	result["うまいもの"] = &Food{
		ID:   1,
		Name: "焼肉",
	}
	result["うまいもの"].Description = "すこぶるうまい"

	return result
}
