package main

import "fmt"

func main() {
	test := make(map[string]interface{})

	test["a"] = true

	b, ok := test["a"].(bool)
	if !ok {
		fmt.Println("ok:", ok)
		fmt.Println("b:", b)
	}
	testFunc(b)
}

func testFunc(a bool) {
	fmt.Println(a)
}
