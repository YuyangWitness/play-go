package main

import "fmt"

type Test struct {
	Key string
}

type Tag struct {
	Key   string
	Value string
	test  Test
}

func main() {
	tag := Tag{
		Key:   "key1",
		Value: "value1",
		test: Test{
			Key: "key2",
		},
	}

	fmt.Print(tag.test.Key)
}
