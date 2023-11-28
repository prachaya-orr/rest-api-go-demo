package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	u := User{
		ID: 1, Name: "John", Age: 20,
	}

	b, err := json.Marshal(u)

	fmt.Printf("type : %T \n", b)
	fmt.Printf("byte : %s \n", b)
	fmt.Println(err)
}
