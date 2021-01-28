package main

import (
	"fmt"
	"play-go/go-sql/operator"
)

// func main() {
// 	DB, err := operator.Start()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// user := operator.User{
// 	Id:       1,
// 	UserName: "yang",
// 	Sex:      "girl",
// 	Location: "shanghai",
// 	Age:      18,
// }

// insertResult, insertErr := operator.InsertUser(DB, &user)
// if insertErr != nil {
// 	fmt.Println(insertErr)
// 	return
// }

// fmt.Printf("%+v is the insert result", insertResult)

// 	result, err := operator.QueryUsers(DB)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Printf("%+v", result)
// }

func main() {
	mySql, err := operator.StartO2()
	defer mySql.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	user := &operator.User2{
		Id:       2,
		Username: "chui",
		Sex:      "girl",
		Location: "shanghai",
		Age:      19,
	}
	operator.CreateUser(mySql, user)
	result := operator.ListUser(mySql)
	fmt.Printf("%+v", result)
}
