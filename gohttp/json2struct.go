package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int
	Name string `json:"nickname"`
	Age  int
}

func main() {
	data := []byte(`{	
		"id":   2,
		"nickname": "Ohm",	
		"age":  26
		}`)

	// u := &User{}
	// err := json.Unmarshal(data, u)

	var u User
	err := json.Unmarshal(data, &u)

	fmt.Printf("% #v\n", u)
	fmt.Println(err)

}
